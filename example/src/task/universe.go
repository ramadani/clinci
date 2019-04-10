package task

import (
	"encoding/json"
	"log"
)

type UniverseTask struct{}

type UniverseData struct {
	Name string
}

func (l *UniverseTask) Key() string {
	return "#"
}

func (l *UniverseTask) Handle(data []byte) error {
	content := &UniverseData{}
	if err := json.Unmarshal(data, content); err != nil {
		log.Print(err)
		return err
	}

	log.Printf("Hello %s from %s", content.Name, l.Key())

	return nil
}
