//go:build wireinject
// +build wireinject

package migration

import (
	"github.com/google/wire"
)

func ProvideProcess() *process {
	panic(wire.Build(NewProcess))
}
