package amqpexporter

import (
	"context"

	"github.com/Azure/go-amqp"
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
	// The following code is for test purposes only... I know, terrible, but I am currently focussing on the receiver side instead of exporter...
	// will work on exporter as soon as receiver is done.
	// the code should be put into a start method and connection should be persisted... again, dev purposes only.

	conn, err := amqp.Dial(context.TODO(), "amqp[s]://<host name of AMQP 1.0 broker>", nil)
	if err != nil {
		// handle error
	}
	session, err := conn.NewSession(context.TODO(), nil)
	if err != nil {
		// handle error
	}

	// create a new sender
	sender, err := session.NewSender(context.TODO(), "<name of peer's receiving terminus>", nil)
	if err != nil {
		// handle error
	}

	// send a message
	err = sender.Send(context.TODO(), amqp.NewMessage([]byte("Hello!")), nil)
	if err != nil {
		// handle error
	}

	return nil
}
