//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"github.com/google/wire"
	"github.com/simon/mengine/infrastructure/engine"
	logger2 "github.com/simon/mengine/infrastructure/logger/wrap"
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
	InitEngine(ctx)
	//rootProvider := InitEngineProvider(ctx)
	//httpProvider := InitHttpProvider(ctx)

	//err := engine.Run(rootProvider, httpProvider)

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
	//if err != nil {
	//	panic(err)
	//}
}

func InitLogger() (*logger2.LoggerWrap, error) {
	panic(wire.Build(logger2.WireLoggerZapSet))
}

func InitEngine(ctx context.Context) engine.Engine {
	panic(wire.Build(engine.WireEngineSet))
	return nil
}

//func InitEngineProvider(ctx context.Context) provider.Provider {
//	panic(wire.Build(engine.WireProviderSet))
//	return nil
//}

//func InitHttpProvider(ctx context.Context) provider.Provider {
//	panic(wire.Build(http.WireProviderSet))
//	return nil
//}
