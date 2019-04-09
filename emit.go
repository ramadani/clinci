package clinci

import "github.com/streadway/amqp"

type Emitter interface {
	Declare(event Event) error
	DeclareAll(events []Event) error
	Dispatcher
}

type rabbitmqEmitter struct {
	ch *amqp.Channel
}

func NewEmitter(ch *amqp.Channel) *rabbitmqEmitter {
	return &rabbitmqEmitter{ch}
}

func (e *rabbitmqEmitter) Declare(event Event) error {
	cog := event.Config()

	err := e.ch.ExchangeDeclare(
		event.Name(),
		event.Kind(),
		cog.Durable,
		cog.AutoDelete,
		cog.Internal,
		cog.NoWait,
		cog.Args,
	)

	if err != nil {
		return err
	}

	return nil
}

func (e *rabbitmqEmitter) DeclareAll(events []Event) error {
	for _, event := range events {
		if err := e.Declare(event); err != nil {
			return err
		}
	}

	return nil
}

func (e *rabbitmqEmitter) Dispatch(pub Publishable) error {
	body, err := pub.Data()
	if err != nil {
		return err
	}

	msg := amqp.Publishing{ContentType: "text/plain", Body: body}
	if err = e.ch.Publish(pub.Name(), pub.Key(), false, false, msg); err != nil {
		return err
	}

	return nil
}
