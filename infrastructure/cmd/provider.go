package cmd

import (
	"context"
	_ "github.com/google/wire"
	provider2 "github.com/simon/mengine/infrastructure/provider"
)

type provider struct {
}

func NewProvider() provider2.Provider {
	return &provider{}
}

func (p provider) Name() string {
	return `cmd`
}

//
//func (p provider) Register() {
//	wire.Build(NewEvent, NewGreeter, NewMessage)
//}

func (p provider) Register(ctx context.Context) error {
	NewCmd()
	return nil
}

func (p provider) Bootstrap(ctx context.Context) error {
	return nil
}
