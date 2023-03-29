//go:build wireinject
// +build wireinject

package a

import (
	"context"
	"github.com/google/wire"
	"github.com/simon/mengine/cmd"
)

func InitializeBaz(ctx context.Context) (cmd.Baz, error) {
	wire.Build(cmd.MegaSet)
	return cmd.Baz{}, nil
}
