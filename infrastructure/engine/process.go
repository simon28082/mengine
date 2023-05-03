package engine

type Process interface {
	Name() string

	// If is Global all process will reload
	//Global() bool
	Dependencies() []string

	Prepare(engine Engine) error

	Shutdown(engine Engine) error
}
