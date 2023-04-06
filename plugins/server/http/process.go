package http

import (
	"context"
	"fmt"
	"github.com/google/wire"
	"github.com/simon/mengine/infrastructure/engine"
)

const (
	ProviderName     = `http`
	ContainerCmdName = `http.engine`
)

var WireProcessSet = wire.NewSet(NewProcess, engine.WireCmdSet)

//var WireProviderSet = wire.NewSet(container.WireContainerSet, cmd2.WireCmdSet, wire.Struct(new(provider), "container", "engine"))

//var WireProviderSet = wire.NewSet(NewProvider)

type process struct {
	ctx     context.Context
	engine  engine.Command
	httpCmd *cmd
}

func NewProcess(ctx context.Context, engine engine.Command) engine.Process {
	return &process{
		ctx:     ctx,
		engine:  engine,
		httpCmd: NewCmd(),
	}
}

func (p *process) Name() string {
	return ProviderName
}

func (p *process) Prepare() error {
	fmt.Println("http provider Register")
	//httpCmd := NewCmd()
	p.httpCmd.Init()
	p.engine.AddCommand(p.httpCmd)
	//cmd, ok := p.container.Get(cmd2.ContainerCmdName)
	//spew.Dump("===ok==", ok)
	//if ok {
	//	fmt.Println("################################")
	//	cmd.(cmd2.Command).AddCommand(httpCmd)
	//}
	//p.container.Put(ContainerCmdName, httpCmd)
	return nil
}

func (p *process) Run() error {
	fmt.Println("http provider Bootstrap")
	//cmd, ok := p.container.Get(ContainerCmdName)
	//if ok {
	//	cmd.(cmd2.Command).Run()
	//}
	p.httpCmd.Run()
	return nil
}

func (p *process) Shutdown() error {
	fmt.Println("http provider shutdown")

	return nil
}
