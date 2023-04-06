package http

import (
	"context"
	"fmt"
	"github.com/google/wire"
	cmd2 "github.com/simon/mengine/infrastructure/engine"
	provider2 "github.com/simon/mengine/infrastructure/provider"
)

const (
	ProviderName     = `http`
	ContainerCmdName = `http.engine`
)

var WireProviderSet = wire.NewSet(NewProvider, cmd2.WireCmdSet)

//var WireProviderSet = wire.NewSet(container.WireContainerSet, cmd2.WireCmdSet, wire.Struct(new(provider), "container", "engine"))

//var WireProviderSet = wire.NewSet(NewProvider)

type provider struct {
	ctx     context.Context
	engine  cmd2.Command
	httpCmd *cmd
}

func NewProvider(ctx context.Context, engine cmd2.Command) provider2.Provider {
	return &provider{
		ctx:     ctx,
		engine:  engine,
		httpCmd: NewCmd(),
	}
}

func (p *provider) Name() string {
	return ProviderName
}

func (p *provider) Prepare() error {
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

func (p *provider) Run() error {
	fmt.Println("http provider Bootstrap")
	//cmd, ok := p.container.Get(ContainerCmdName)
	//if ok {
	//	cmd.(cmd2.Command).Run()
	//}
	p.httpCmd.Run()
	return nil
}

func (p *provider) Shutdown() error {
	fmt.Println("http provider shutdown")

	return nil
}
