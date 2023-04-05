//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"github.com/google/wire"
	"github.com/simon/mengine/infrastructure/container"
	"github.com/simon/mengine/infrastructure/engine"
	"github.com/simon/mengine/infrastructure/provider"
	"github.com/simon/mengine/plugins/server/http"
)

func main() {
	//var (
	//	cmdFunc = func(next chain.ChainFunc) chain.ChainFunc {
	//		InitPro
	//		return next
	//	}
	//	httpFunc = func(next chain.ChainFunc) chain.ChainFunc {
	//		http.NewProvider()
	//		return next
	//	}
	//)

	//chain.BuildChain(cmdFunc, httpFunc)

	ctx := context.Background()
	rootProvider := InitEngineProvider(ctx)
	httpProvider := InitHttpProvider(ctx)

	err := engine.Run(rootProvider, httpProvider)

	//err := chain.BuildChain(func(next chain.HandlerFunc) (chain.HandlerFunc, error) {
	//	println("################################################################")
	//	rootProvider.Register()
	//	httpProvider.Register()
	//	return next, nil
	//}, func(next chain.HandlerFunc) (chain.HandlerFunc, error) {
	//	rootProvider.Bootstrap()
	//	httpProvider.Bootstrap()
	//	return next, nil
	//}, func(next chain.HandlerFunc) (chain.HandlerFunc, error) {
	//	httpProvider.Shutdown()
	//	rootProvider.Shutdown()
	//	return next, nil
	//})
	if err != nil {
		panic(err)
	}
}

func InitEngineProvider(ctx context.Context) provider.Provider {
	panic(wire.Build(container.WireContainerSet, engine.WireProviderSet, engine.WireCmdSet))
	return nil
}

func InitHttpProvider(ctx context.Context) provider.Provider {
	panic(wire.Build(container.WireContainerSet, engine.WireCmdSet, http.WireProviderSet))
	return nil
}
