package clinci

type defaultListener struct {
	queuer   Queuer
	consumer Consumer
	task     Task
}

type defaultQueuer struct {
	name string
}

type defaultConsumer struct{}

func CreateFromDefaultListener(task Task) Listener {
	return &defaultListener{
		queuer:   DefaultQueuer(),
		consumer: DefaultConsumer(),
		task:     task,
	}
}

func DefaultQueuer() Queuer {
	return &defaultQueuer{}
}

func DefaultConsumer() Consumer {
	return &defaultConsumer{}
}

func (l *defaultListener) Queuer() Queuer {
	return l.queuer
}

func (l *defaultListener) Consumer() Consumer {
	return l.consumer
}

func (l *defaultListener) Key() string {
	return l.task.Key()
}

func (l *defaultListener) Handle(data []byte) error {
	return l.task.Handle(data)
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

func (c *defaultConsumer) Name() string {
	return ""
}

func (c *defaultConsumer) Config() *Config {
	return &Config{
		AutoAck: true,
		NoLocal: false,
		NoWait:  false,
		Args:    nil,
	}
}
