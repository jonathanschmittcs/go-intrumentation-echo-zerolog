package myserver

import (
	"os"
	"strings"

	"github.com/jonathanschmittcs/go-intrumentation-echo-zerolog/internal/mylogger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func myCustomRequestMiddleware() echo.MiddlewareFunc {
	logHeaders = os.Getenv("LOG_HEADERS") == "true"
	lohHeadersNames = strings.Split(os.Getenv("LOG_HEADERS_NAMES"), ",")

	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:        true,
		LogStatus:     true,
		LogRemoteIP:   true,
		LogHost:       true,
		LogMethod:     true,
		LogError:      true,
		LogLatency:    true,
		LogUserAgent:  true,
		LogHeaders:    lohHeadersNames,
		LogValuesFunc: LogLogValuesFunc,
	})
}

func LogLogValuesFunc(c echo.Context, v middleware.RequestLoggerValues) error {
	var errorMsg string
	if v.Error != nil {
		errorMsg = v.Error.Error()
	}

	event := mylogger.Info(c.Request().Context()).
		Str("uri", v.URI).
		Int("status", v.Status).
		Str("remote_ip", v.RemoteIP).
		Str("host", v.Host).
		Str("method", v.Method).
		Str("error", errorMsg).
		Str("latency", v.Latency.String()).
		Str("user_agent", v.UserAgent)

	if logHeaders && len(lohHeadersNames) > 0 {
		event.Interface("headers", v.Headers)
	}

	event.Timestamp().Msg("http_request")
	return nil
}
