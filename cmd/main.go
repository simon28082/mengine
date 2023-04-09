package main

import (
	"context"
	"github.com/simon/mengine/infrastructure/engine"
	"github.com/simon/mengine/plugins/server/http"
	"github.com/simon/mengine/plugins/server/tcp"
	"os"
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

	//ctx := context.Background()
	//InitEngine(ctx)
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

	// Create a new logger
	//logger := NewLogger()
	//
	//// Create a set of providers that includes the logger and the MyService constructor
	//MySet := wire.NewSet(
	//	wire.Value(logger),
	//	MyServiceFunc,
	//)

	// Use Wire to generate the dependencies for MyService

	//provider()

	//logger := InitLogger()
	//
	//config, err := config2.NewConfig(source.NewFile(os2.RunPath("1.json")))
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//
	//engine := InitEngine(context.Background(), logger, config)
	//
	//engine.Mount(http.NewProcess(engine))
	//
	//engine.Run()

	e := engine.EngineProvide(context.Background())

	//e := engine.NewEngine(context.Background())
	//
	e.Mount(http.NewProcess(), tcp.NewProcess())

	e.Run(os.Args...)

	//provideSource()

	//s := ""
	//wire.Build(
	//	wire.Value(s), // inject empty string
	//	provideString, // provide string value
	//)

	//InitConfig("abc")
	//provideSource("path/to/file")

}

//var logger2 = NewLogger()
//
//func provider() *MyService {
//	MySet := wire.NewSet(
//
//		wire.Value(logger2),
//	)
//	wire.Build(MyServiceFunc, MySet)
//	return &MyService{}
//}

//func main() {
//	MySet := wire.NewSet(
//		provideComplexStruct,
//		injectMyService,
//	)
//
//	wire.Build(MySet)
//}

//func InitLogger() *logger22.LoggerWrap {
//	panic(wire.Build(logger22.WireLoggerSet))
//}

// var wireValue = wire.Value(logger2.C)
//
//	func providerLoggerValue() *logger2.LoggerWrap {
//		panic(wire.Build(wireValue))
//	}
//func providePath(path string) string {
//	return path
//}
//
//var pathSet = wire.NewSet(
//	providePath,
//)

//	func provide(path string) source.Source {
//		panic(wire.Build(source.NewFile, wire.NewSet(providePath)))
//	}

//func provideString() string {
//	return "hello, world"
//}

//func pb(p string) source.PathString {
//	return source.PathString(p)
//}

//func InitLogger() *wrap.LoggerWrap {
//	panic(wire.Build(wrap.WireLoggerZapDevelopmentSet))
//}

//func providerS() source.PathString {
//	return "abc"
//}
//
//func provideSource(path string) source.Source {
//	wire.Build(wire.NewSet(
//		source.NewFile,
//		providerS,
//	))
//	return nil
//}

//func InitConfig(path string) config.Config {
//	panic(wire.Build(config.NewConfig, provide))
//}

//func InitEngine(ctx context.Context, logger *wrap.LoggerWrap, config config2.Config) engine.Engine {
//	panic(wire.Build(engine.WireEngineSet))
//	return nil
//}

//
//func InitHttpProcess(ctx context.Context) engine.Process {
//	panic(wire.Build(http.WireProcessSet))
//}

//func InitEngineProvider(ctx context.Context) provider.Provider {
//	panic(wire.Build(engine.WireProviderSet))
//	return nil
//}

//func InitHttpProvider(ctx context.Context) provider.Provider {
//	panic(wire.Build(http.WireProviderSet))
//	return nil
//}
