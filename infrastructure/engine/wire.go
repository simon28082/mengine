//go:build wireinject
// +build wireinject

package engine

import (
	"context"
	"github.com/google/wire"
)

func ProvideEngine(ctx context.Context) Engine {
	panic(wire.Build(WireEngineSet))
}
