package engine

type Process interface {
	Name() string

	Global() bool

	Prepare(engine Engine) error

	Shutdown(engine Engine) error
}
