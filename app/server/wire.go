//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"
	"github.com/simon/mengine/infrastructure/engine"
)

func ProvideProcess() engine.Process {
	panic(wire.Build(WireHttpProcessSet))
}
