package main

import (
	"log"

	"github.com/ramadani/clinci"
	"github.com/ramadani/clinci/example/src/event"
	"github.com/ramadani/clinci/example/src/task"
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

	worker := clinci.NewWorker(ch)
	worker.Declare(&clinci.EventListener{
		Event: &event.HelloEvent{},
		Listeners: []clinci.Listener{
			clinci.CreateFromDefaultListener(&task.WorldTask{}),
			clinci.CreateFromDefaultListener(&task.SystemTask{}),
			clinci.CreateFromDefaultListener(&task.UniverseTask{}),
		},
	})

	listen, err := worker.Listen()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Waiting for logs. To exit press CTRL+C")

	<-listen
}
