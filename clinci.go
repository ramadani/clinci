package clinci

import "github.com/streadway/amqp"

// Dispatcher for publisher
type Dispatcher interface {
	Dispatch(pub Publishable) error
}

// Event
type Event interface {
	Configurable
	Kind() string
}

type Listener interface {
	Task
	Queuer() Queuer
	Consumer() Consumer
}

type Publishable interface {
	Name() string
	Data() ([]byte, error)
	Routing
}

type Queuer interface {
	SetName(name string)
	Configurable
}

type Consumer interface {
	Configurable
}

type Configurable interface {
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
	AutoAck    bool
	Internal   bool
	NoLocal    bool
	NoWait     bool
	Args       amqp.Table
}
