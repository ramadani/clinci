package clinci

type Routing interface {
	Key() string
}

type Dispatcher interface {
	Dispatch(pub Publishable) error
}
