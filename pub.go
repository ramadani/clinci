package clinci

type Publishable interface {
	Name() string
	Data() ([]byte, error)
	Routing
}
