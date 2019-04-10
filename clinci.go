package clinci

import "github.com/streadway/amqp"

type Dispatcher interface {
	Dispatch(pub Publishable) error
}

type Event interface {
	Declarable
	Kind() string
}

type Listener interface {
	Queuer() Queuer
	Routing
	Task
}

type Publishable interface {
	Name() string
	Data() ([]byte, error)
	Routing
}

type Queuer interface {
	SetName(name string)
	Declarable
}

type Declarable interface {
	Name() string
	Config() *Config
}

type Task interface {
	Handle(data []byte) error
}

type Routing interface {
	Key() string
}

type Config struct {
	Durable    bool
	AutoDelete bool
	Internal   bool
	NoWait     bool
	Args       amqp.Table
}

type defaultQueuer struct {
	name string
}

func DefaultQueuer() Queuer {
	return &defaultQueuer{}
}

func (q *defaultQueuer) SetName(name string) {
	q.name = name
}

func (q *defaultQueuer) Name() string {
	return q.name
}

func (q *defaultQueuer) Config() *Config {
	return &Config{
		Durable:    false,
		AutoDelete: false,
		NoWait:     false,
		Args:       nil,
	}
}
