package mytracer

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/jonathanschmittcs/go-intrumentation-echo-zerolog/internal/mylogger"
	"go.elastic.co/apm/transport"
	"go.elastic.co/apm/v2"
)

var (
	Tracer            = apm.DefaultTracer()
	ApmEnabled        bool
	ApmServiceName    string
	ApmServiceVersion string
)

func Init() error {
	ApmEnabled = os.Getenv("ELASTIC_APM_ACTIVE") == "true"
	if !ApmEnabled {
		return nil
	}

	ApmServiceName = os.Getenv("ELASTIC_APM_SERVICE_NAME")
	ApmServiceVersion = os.Getenv("ELASTIC_APM_SERVICE_VERSION")

	var err error
	Tracer, err = apm.NewTracer(fmt.Sprint(ApmServiceName), ApmServiceVersion)
	if err != nil {
		return err
	}

	Tracer.SetSanitizedFieldNames(strings.Split(os.Getenv("ELASTIC_APM_SANITIZE_FIELD_NAMES"), ",")...)

	_, err = transport.InitDefault()
	if err != nil {
		return err
	}

	mylogger.Info(context.Background()).Msg("APM Tracer inicializado com sucesso")

	return nil
}
