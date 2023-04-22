//go:build wireinject
// +build wireinject

package engine

import (
	"context"
	"github.com/google/wire"
	"github.com/simon28082/mengine/infrastructure/logger"
	logger2 "github.com/simon28082/mengine/infrastructure/logger/wrap"
	"github.com/simon28082/mengine/infrastructure/logger/zap"
)

func ProvideEngine(ctx context.Context) Engine {
	panic(wire.Build(WireEngineSet))
}

func ProvideZapDevLogger() logger.Wrap {
	return logger2.NewLogger(zap.NewZapDevelopment())
}

func ProvideZapProdLogger() logger.Wrap {
	return logger2.NewLogger(zap.NewZapProduction())
}
