package lacia

type LInterface interface {
	New()
}

type ErrInterface interface {
	error
	Errors() []error
}
