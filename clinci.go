package clinci

import "github.com/streadway/amqp"

type Event interface {
	Declarable
	Kind() string
}

type Dispatcher interface {
	Dispatch(pub Publishable) error
}

type Publishable interface {
	Name() string
	Data() ([]byte, error)
	Routing
}

type Queue interface {
	SetName(name string)
	Declarable
}

type Declarable interface {
	Name() string
	Config() *Config
}

type Task interface {
	Routing
	Handle(data []byte) error
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
