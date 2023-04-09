//go:build wireinject
// +build wireinject

package engine

import (
	"context"
	"github.com/google/wire"
	logger2 "github.com/simon/mengine/infrastructure/logger/wrap"
)

func EngineProvide(ctx context.Context) Engine {
	panic(wire.Build(WireEngineSet))
}

func LoggerProvide() *logger2.LoggerWrap {
	panic(wire.Build(logger2.WireLoggerZapDevelopmentSet))
}
