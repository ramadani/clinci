package clinci

type Publishable interface {
	Name() string
	Kind() string
	Data() ([]byte, error)
	Routing
}
