package myserver

import (
	"github.com/jonathanschmittcs/go-intrumentation-echo-zerolog/internal/mylogger"
	"github.com/jonathanschmittcs/go-intrumentation-echo-zerolog/internal/mytracer"
	"github.com/labstack/echo/v4"
	"go.elastic.co/apm/module/apmechov4/v2"
)

var (
	logHeaders      bool     = false
	lohHeadersNames []string = []string{}
)

func Start(port string) {
	e := echo.New()
	e.Use(myCustomRequestMiddleware())

	if mytracer.Tracer != nil {
		e.Use(apmechov4.Middleware(apmechov4.WithTracer(mytracer.Tracer)))
	}

	e.GET("/health", func(c echo.Context) error {
		mylogger.Info(c.Request().Context()).Msg("Health check")
		return c.String(200, "OK")
	})

	e.POST("/message", func(c echo.Context) error {
		ctx := c.Request().Context()
		mylogger.Info(ctx).Msg("Message received")

		var body map[string]interface{}
		if err := c.Bind(&body); err != nil {
			mylogger.Error(ctx, err).Msg("Error binding request body")
			return c.JSON(400, err)
		}

		return c.JSON(200, map[string]string{"message": "Hello, World!"})
	})

	e.Start(port)
}
