package clinci

import "github.com/streadway/amqp"

type RabbitMQ struct {
	ch *amqp.Channel
}

func NewRabbitMQ(ch *amqp.Channel) *RabbitMQ {
	return &RabbitMQ{ch}
}

func (d *RabbitMQ) Dispatch(pub Publishable) error {
	body, err := pub.Data()
	if err != nil {
		return err
	}

	msg := amqp.Publishing{ContentType: "text/plain", Body: body}
	if err = d.ch.Publish(pub.Name(), pub.Key(), false, false, msg); err != nil {
		return err
	}

	return nil
}
