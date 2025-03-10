package amqpexporter

import (
	"context"

	"go.opentelemetry.io/collector/pdata/plog"
)

type amqpLogExporter struct {
	conf Config
}

func newLogsExporter(config Config) *amqpLogExporter {
	return &amqpLogExporter{
		conf: config,
	}
}

func (e *amqpLogExporter) pushLogs(
	ctx context.Context,
	ld plog.Logs,
) error {

	return nil
}
