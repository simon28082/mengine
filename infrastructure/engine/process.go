package engine

type Process interface {
	Name() string

	Prepare() error

	Run() error

	Shutdown() error
}
