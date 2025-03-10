package amqpexporter

import (
	"context"
	"errors"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/amqpexporter/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/exporter/xexporter"
)

var (
	typeStr = component.MustNewType("amqp")
)

var errInvalidConfig = errors.New("config was not an AMQP exporter config")

type logsExporter struct {
	cfg *Config
}

// NewFactory creates a Datadog exporter factory
func NewFactory() exporter.Factory {
	return xexporter.NewFactory(
		metadata.Type,
		createDefaultConfig,
		xexporter.WithLogs(createLogsExporter, metadata.LogsStability),
	)
}

func createDefaultConfig() component.Config {
	c := &Config{
		Address: "",
	}
	return c
}

func createLogsExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exporter.Logs, error) {
	oCfg := *(cfg.(*Config))
	le := newLogsExporter(oCfg)

	return exporterhelper.NewLogs(
		ctx,
		set,
		cfg,
		le.pushLogs,
	)
}
