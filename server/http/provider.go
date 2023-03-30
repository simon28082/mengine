package http

import (
	"context"
	provider2 "github.com/simon/mengine/infrastructure/provider"
)

type provider struct {
}

func NewProvider() provider2.Provider {
	return &provider{}
}

func (p provider) Name() string {
	return `http.server`
}

func (p provider) Register(ctx context.Context) error {
	return nil
}

func (p provider) Bootstrap(ctx context.Context) error {
	return nil
}
