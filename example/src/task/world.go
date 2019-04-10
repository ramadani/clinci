package task

import (
	"encoding/json"
	"log"
)

type WorldTask struct{}

type WorldData struct {
	Name string
}

func (l *WorldTask) Key() string {
	return "clinci.example.hello.world"
}

func (l *WorldTask) Handle(data []byte) error {
	content := &WorldData{}
	if err := json.Unmarshal(data, content); err != nil {
		log.Print(err)
		return err
	}

	log.Printf("Hello %s from %s", content.Name, l.Key())

	return nil
}
