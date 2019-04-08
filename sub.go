package clinci

type Subscribable interface {
	Routing
	Handle(data []byte) error
}
