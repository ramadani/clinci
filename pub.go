package clinci

type Publishable interface {
	Name() string
	Key() string
	Data() ([]byte, error)
}
