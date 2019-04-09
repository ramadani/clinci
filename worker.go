package clinci

import (
	"github.com/streadway/amqp"
)

type Worker interface {
	Declare(eventTask EventTask) error
	DeclareAll(eventTasks []EventTask) error
}

type EventTask struct {
	Event Event
	Queue Queue
	Tasks []Task
}

type rabbitmqWorker struct {
	ch *amqp.Channel
}

func (w *rabbitmqWorker) Declare(eventTask EventTask) error {
	// Declaring the Exchange
	event := eventTask.Event
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

	// Declaring the Queue
	queue := eventTask.Queue
	qCog := queue.Config()

	q, err := w.ch.QueueDeclare("", qCog.Durable, qCog.AutoDelete, qCog.Exclusive,
		qCog.NoWait, qCog.Args)

	if err != nil {
		return err
	}

	queue.SetName(q.Name)

	// Binding the Queue
	tasks := eventTask.Tasks
	for _, task := range tasks {
		err = w.ch.QueueBind(queue.Name(), task.Key(), event.Name(), qCog.NoWait, qCog.Args)
		if err != nil {
			return err
		}
	}

	return nil
}

func (w *rabbitmqWorker) DeclareAll(eventTasks []EventTask) error {
	for _, eventTask := range eventTasks {
		if err := w.Declare(eventTask); err != nil {
			return err
		}
	}

	return nil
}
