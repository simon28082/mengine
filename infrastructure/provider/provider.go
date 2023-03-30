package provider

import "context"

type Provider interface {
	Name() string

	Register(ctx context.Context) error

	Bootstrap(ctx context.Context) error
}
