package clinci

import (
	"github.com/streadway/amqp"
)

type Worker interface {
	Declare(el *EventListener) error
	DeclareAll(eventListeners []*EventListener) error
	Listen() error
}

type EventListener struct {
	Event     Event
	Listeners []Listener
}

type rabbitmqWorker struct {
	ch             *amqp.Channel
	eventListeners []*EventListener
}

func NewWorker(ch *amqp.Channel) *rabbitmqWorker {
	eventListeners := make([]*EventListener, 0)

	return &rabbitmqWorker{ch, eventListeners}
}

func (w *rabbitmqWorker) Declare(el *EventListener) error {
	// Declaring the Exchange
	event := el.Event
	eCog := event.Config()

	err := w.ch.ExchangeDeclare(
		event.Name(),
		event.Kind(),
		eCog.Durable,
		eCog.AutoDelete,
		eCog.Internal,
		eCog.NoWait,
		eCog.Args,
	)

	if err != nil {
		return err
	}

	for _, lis := range el.Listeners {
		queuer := lis.Queuer()
		qCog := queuer.Config()

		// Declaring the Queue
		q, err := w.ch.QueueDeclare(
			"",
			qCog.Durable,
			qCog.AutoDelete,
			true,
			qCog.NoWait,
			qCog.Args,
		)

		if err != nil {
			return err
		}

		queuer.SetName(q.Name)

		// Binding the Queue
		err = w.ch.QueueBind(queuer.Name(), lis.Key(), event.Name(), qCog.NoWait, qCog.Args)
		if err != nil {
			return err
		}
	}

	w.push(el)

	return nil
}

func (w *rabbitmqWorker) DeclareAll(eventListeners []*EventListener) error {
	for _, el := range eventListeners {
		if err := w.Declare(el); err != nil {
			return err
		}
	}

	return nil
}

func (w *rabbitmqWorker) Listen() (<-chan bool, error) {
	listen := make(chan bool)

	for _, el := range w.eventListeners {
		for _, lis := range el.Listeners {
			msgs, err := w.ch.Consume(
				lis.Queuer().Name(),
				"",
				true,
				false,
				false,
				false,
				nil,
			)

			if err != nil {
				return listen, err
			}

			go w.consume(msgs, lis)
		}
	}

	return listen, nil
}

func (w *rabbitmqWorker) push(el *EventListener) {
	w.eventListeners = append(w.eventListeners, el)
}

func (w *rabbitmqWorker) consume(msgs <-chan amqp.Delivery, task Task) {
	for d := range msgs {
		task.Handle(d.Body)
	}
}
