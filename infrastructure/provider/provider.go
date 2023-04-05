package provider

type Provider interface {
	Name() string

	Prepare() error

	Run() error

	Shutdown() error
}
