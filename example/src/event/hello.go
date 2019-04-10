package event

import (
	"encoding/json"

	"github.com/ramadani/clinci"
)

type HelloEvent struct {
	data       *HelloData
	routingKey string
}

type HelloData struct {
	Name string
}

func NewHelloEvent(name string, key string) *HelloEvent {
	data := &HelloData{Name: name}

	return &HelloEvent{data: data, routingKey: key}
}

func (e *HelloEvent) Name() string {
	return "clinci.hello"
}

func (e *HelloEvent) Kind() string {
	return "topic"
}

func (e *HelloEvent) Data() ([]byte, error) {
	return json.Marshal(e.data)
}

func (a *HelloEvent) Key() string {
	return a.routingKey
}

func (a *HelloEvent) Config() *clinci.Config {
	return &clinci.Config{
		Durable:    true,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
		Args:       nil,
	}
}
