package tcp

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/simon28082/mengine/infrastructure/engine"
	"github.com/spf13/cobra"
)

type process struct {
	ctx context.Context
}

func NewProcess() engine.Process {
	return &process{}
}

func (p *process) Name() string {
	return `tcp`
}

func (m *process) Dependencies() []string {
	return nil
}

func (p *process) Cobra() *cobra.Command {
	cli := &cobra.Command{
		Use: `tcp`,
	}
	cli.Run = func(cmd *cobra.Command, args []string) {
		fmt.Println("tcp start..........")
		fmt.Println("tcp==================")
		spew.Dump(cmd.Flag("log-path").Value.String())
	}

	return cli

}
func (p *process) Global() bool {
	return false
}
func (p *process) Prepare(e engine.Engine) error {
	fmt.Println("tcp provider Register")
	return nil
}

func (p *process) Shutdown(e engine.Engine) error {
	fmt.Println("tcp provider shutdown")

	return nil
}
