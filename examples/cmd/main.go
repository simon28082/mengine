package main

import (
	"context"
	"fmt"
	"github.com/simon28082/mengine/infrastructure/engine"
	"github.com/simon28082/mengine/infrastructure/logger"
	"github.com/simon28082/mengine/plugins/server/http"
	"os"
)

func main() {
	e := engine.ProvideEngine(context.Background())

	e.Mount(http.NewProcess())

	if err := e.Run(os.Args...); err != nil {
		fmt.Println(err)
		//os.Exit(1)
		logger.Error(err.Error())
	} else {
		logger.Error("abc")

	}

}
