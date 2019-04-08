package clinci

import "github.com/streadway/amqp"

type Dispatcher interface {
	Dispatch(pub Publishable) error
}

type Event interface {
	Declarable
	Kind() string
}

type Declarable interface {
	Name() string
	Config() *Config
}

type Routing interface {
	Key() string
}

type Config struct {
	Durable    bool
	AutoDelete bool
	Internal   bool
	Exclusive  bool
	NoWait     bool
	Args       amqp.Table
}
