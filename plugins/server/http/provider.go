package http

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/wire"
	"github.com/simon/mengine/infrastructure/container"
	cmd2 "github.com/simon/mengine/infrastructure/engine"
	provider2 "github.com/simon/mengine/infrastructure/provider"
)

const (
	ProviderName     = `http`
	ContainerCmdName = `http.engine`
)

var WireProviderSet = wire.NewSet(NewProvider)

type provider struct {
	ctx       context.Context
	container container.Container
}

func NewProvider(ctx context.Context, container container.Container) provider2.Provider {
	return &provider{
		ctx:       ctx,
		container: container,
	}
}

func (p *provider) Name() string {
	return ProviderName
}

func (p *provider) Prepare() error {
	fmt.Println("http provider Register")
	httpCmd := NewCmd()
	httpCmd.Init()
	cmd, ok := p.container.Get(cmd2.ContainerCmdName)
	spew.Dump("===ok==", ok)
	if ok {
		fmt.Println("################################")
		cmd.(cmd2.Command).AddCommand(httpCmd)
	}
	p.container.Put(ContainerCmdName, httpCmd)
	return nil
}

func (p *provider) Run() error {
	fmt.Println("http provider Bootstrap")
	cmd, ok := p.container.Get(ContainerCmdName)
	if ok {
		cmd.(cmd2.Command).Run()
	}
	return nil
}

func (p *provider) Shutdown() error {
	fmt.Println("http provider shutdown")

	return nil
}
