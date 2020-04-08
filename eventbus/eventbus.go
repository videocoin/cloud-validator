package eventbus

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	v1 "github.com/videocoin/cloud-api/validator/v1"
	"github.com/videocoin/cloud-pkg/mqmux"
)

type EventBus struct {
	logger *logrus.Entry
	uri    string
	name   string
	mq     *mqmux.WorkerMux
}

func NewEventBus(uri string, opts ...Option) (*EventBus, error) {
	eb := &EventBus{
		uri: uri,
	}
	for _, o := range opts {
		if err := o(eb); err != nil {
			return nil, err
		}
	}

	mq, err := mqmux.NewWorkerMux(eb.uri, eb.name)
	if err != nil {
		return nil, err
	}

	eb.mq = mq

	return eb, nil
}

func (e *EventBus) Start() error {
	err := e.mq.Publisher("validator.events")
	if err != nil {
		return err
	}

	return e.mq.Run()
}

func (e *EventBus) Stop() error {
	return e.mq.Close()
}

func (e *EventBus) EmitEvent(ctx context.Context, event *v1.Event) error {
	headers := make(amqp.Table)

	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		ext.SpanKindRPCServer.Set(span)
		ext.Component.Set(span, "transcoder")
		err := span.Tracer().Inject(
			span.Context(),
			opentracing.TextMap,
			mqmux.RMQHeaderCarrier(headers),
		)
		if err != nil {
			e.logger.Errorf("failed to span inject: %s", err)
		}
	}

	err := e.mq.PublishX("validator.events", event, headers)
	if err != nil {
		return err
	}

	return nil
}
