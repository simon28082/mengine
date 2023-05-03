package http

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/wire"
	"github.com/simon28082/mengine/infrastructure/engine"
	logger2 "github.com/simon28082/mengine/infrastructure/logger"
	"github.com/spf13/cobra"
)

// var WireProcessSet = wire.NewSet(NewProcess, engine.WireCmdSet, logger2.WireZapLoggerProductionSet)
var WireProcessSet = wire.NewSet(NewProcess)

//var WireProviderSet = wire.NewSet(container.WireContainerSet, cmd2.WireCmdSet, wire.Struct(new(provider), "container", "engine"))

//var WireProviderSet = wire.NewSet(NewProvider)

type process struct {
	ctx    context.Context
	logger *logger2.WrapLogger
	engine engine.Engine
}

func NewProcess() engine.Process {
	return &process{}
}

func (p *process) Global() bool {
	return false
}
func (p *process) Name() string {
	return `http`
}

func (m *process) Dependencies() []string {
	return nil
}

func (p *process) Cobra() *cobra.Command {
	cli := &cobra.Command{
		Use: `http`,
	}
	cli.Run = func(cmd *cobra.Command, args []string) {
		fmt.Println("http start..........")
		fmt.Println("http==================")
		spew.Dump(cmd.Flag("log-path").Value.String())
	}

	return cli
}

func (p *process) Prepare(e engine.Engine) error {
	fmt.Println("http provider Register")
	//p.httpCmd.Init()
	//rootCommand, _ := p.engine.Get(`command.root`)
	//rootCommand.(engine.Command).AddCommand(p.httpCmd)

	//httpCmd := NewCmd()

	//p.engine.AddCommand(p.httpCmd)
	//cmd, ok := p.container.Get(cmd2.ContainerCmdName)
	//spew.Dump("===ok==", ok)
	//if ok {
	//	fmt.Println("################################")
	//	cmd.(cmd2.Command).AddCommand(httpCmd)
	//}
	//p.container.Put(ContainerCmdName, httpCmd)
	return nil
}

func (p *process) Shutdown(e engine.Engine) error {
	fmt.Println("http provider shutdown")

	return nil
}
