package clinci

type Emitter interface {
	Declare(event Event) error
	DeclareAll(events []Event) error
	Dispatcher
}
