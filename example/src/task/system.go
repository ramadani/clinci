package task

import (
	"encoding/json"
	"log"
)

type SystemTask struct{}

type SystemData struct {
	Name string
}

func (l *SystemTask) Key() string {
	return "clinci.example.hi.*"
}

func (l *SystemTask) Handle(data []byte) error {
	content := &SystemData{}
	if err := json.Unmarshal(data, content); err != nil {
		log.Print(err)
		return err
	}

	log.Printf("Hello %s from %s", content.Name, l.Key())

	return nil
}
