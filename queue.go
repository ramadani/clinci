package clinci

type Queueable interface {
	Queuer() Queuer
}

type Queuer interface {
	SetName(name string)
	Declarable
}
