package clinci

type Subscribable interface {
	Handle(data []byte) error
}
