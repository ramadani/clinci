package clinci

import "github.com/streadway/amqp"

type Dispatcher interface {
	Dispatch(pub Publishable) error
}

type Routing interface {
	Key() string
}

type Declarable interface {
	Config() *Config
}

type Config struct {
	Durable    bool
	AutoDelete bool
	Internal   bool
	Exclusive  bool
	NoWait     bool
	Args       amqp.Table
}
