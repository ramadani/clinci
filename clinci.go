package clinci

type Dispatcher interface {
	Dispatch(pub Publishable) error
}
