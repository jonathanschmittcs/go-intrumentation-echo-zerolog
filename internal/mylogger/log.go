package mylogger

import (
	"context"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.elastic.co/apm/module/apmzerolog"
	"go.elastic.co/apm/v2"
)

var ApmServiceName string

func Init() {
	ApmServiceName = os.Getenv("ELASTIC_APM_SERVICE_NAME")
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = apmzerolog.MarshalErrorStack
	log.Logger.Info().Timestamp().Msg("Logger initialized successfully")
}

func Info(ctx context.Context) *zerolog.Event {
	event := log.Info().Ctx(ctx).Timestamp().Str("service_name", ApmServiceName)
	tx := apm.TransactionFromContext(ctx)
	if tx != nil {
		event.Str("trace.id", tx.TraceContext().Trace.String())
		event.Str("transaction.id", tx.TraceContext().Span.String())
		event.Str("span.id", tx.TraceContext().Span.String())
	}
	return event
}

func Error(ctx context.Context, err error) *zerolog.Event {
	event := log.Error().Ctx(ctx).Timestamp().Str("service_name", ApmServiceName).Err(err)
	tx := apm.TransactionFromContext(ctx)
	if tx != nil {
		event.Str("trace.id", tx.TraceContext().Trace.String())
		event.Str("transaction.id", tx.TraceContext().Span.String())
		event.Str("span.id", tx.TraceContext().Span.String())
	}
	return event
}
