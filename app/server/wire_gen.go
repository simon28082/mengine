// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package server

import (
	"github.com/simon/mengine/infrastructure/engine"
)

// Injectors from wire.go:

func ProvideProcess() engine.Process {
	engineProcess := NewProcess()
	return engineProcess
}
