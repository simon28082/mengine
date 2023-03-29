package provider

type Provider interface {
	Name() string

	Register()
	
	Boot()
}
