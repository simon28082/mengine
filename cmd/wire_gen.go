// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/simon/mengine/infrastructure/engine"
	"github.com/simon/mengine/infrastructure/provider"
	"github.com/simon/mengine/plugins/server/http"
)

// Injectors from main.go:

func InitEngineProvider(ctx context.Context) provider.Provider {
	command := _wireCmdValue
	providerProvider := engine.NewProvider(ctx, command)
	return providerProvider
}

var (
	_wireCmdValue = engine.NewCmd()
)

func InitHttpProvider(ctx context.Context) provider.Provider {
	command := _wireCmdValue
	providerProvider := http.NewProvider(ctx, command)
	return providerProvider
}

// main.go:

func main() {

	ctx := context.Background()
	rootProvider := InitEngineProvider(ctx)
	httpProvider := InitHttpProvider(ctx)

	err := engine.Run(rootProvider, httpProvider)

	if err != nil {
		panic(err)
	}
}
