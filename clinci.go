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

type Subscribable interface {
	Routing
	Handle(data []byte) error
}

type Queueable interface {
	Queuer() Queuer
}

type Queuer interface {
	SetName(name string)
	Declarable
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
