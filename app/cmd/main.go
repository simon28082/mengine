package main

import (
	"context"
	"fmt"
	"github.com/simon/mengine/app/server"
	"github.com/simon/mengine/infrastructure/engine"
	"os"
)

func main() {
	e := engine.EngineProvide(context.Background())

	e.Mount(server.ProvideProcess())

	if err := e.Run(os.Args...); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
