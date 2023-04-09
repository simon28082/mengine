package main

import (
	"context"
	"fmt"
	"github.com/simon28082/mengine/infrastructure/engine"
	"github.com/simon28082/mengine/plugins/server/http"
	"os"
)

func main() {
	e := engine.EngineProvide(context.Background())

	e.Mount(http.NewProcess())

	if err := e.Run(os.Args...); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
