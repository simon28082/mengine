package main

import (
	"context"
	"fmt"
	"github.com/simon28082/mengine/infrastructure/engine"
	"github.com/urfave/cli/v3"
	"log"
)

func main() {
	eng := engine.NewEngine(
		engine.WithContext(context.TODO()),
		engine.WithAction(func(ctx context.Context, command *cli.Command) error {
			fmt.Println("running")
			return nil
		}),
	)
	if err := eng.Run(); err != nil {
		log.Fatal(err)
	}
}
