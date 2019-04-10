package main

import (
	"log"

	"github.com/ramadani/clinci"
	"github.com/ramadani/clinci/example/src/event"
	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	emitter := clinci.NewEmitter(ch)
	emitter.Declare(&event.HelloEvent{})

	type data struct {
		key, name string
	}

	listOfData := []data{
		data{key: "clinci.example.hello.world", name: "Earth"},
		data{key: "clinci.example.hi.solar", name: "Solar System"},
		data{key: "clinci.example.milky", name: "Milky Way"},
		data{key: "clinci.example.andro", name: "Andromeda"},
		data{key: "clinci.example.andro", name: "Andromeda"},
	}

	for _, d := range listOfData {
		helloEvent := event.NewHelloEvent(d.name, d.key)

		if err := emitter.Dispatch(helloEvent); err != nil {
			log.Fatal(err)
		} else {
			log.Print("Send by key ", helloEvent.Key())
		}
	}
}
