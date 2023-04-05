package engine

import (
	"context"
	"fmt"
	"github.com/google/wire"
	"github.com/simon/mengine/infrastructure/container"
	provider2 "github.com/simon/mengine/infrastructure/provider"
)

const (
	ProviderName     = `root`
	ContainerCmdName = `root.engine`
)

var WireProviderSet = wire.NewSet(NewProvider)

type provider struct {
	ctx       context.Context
	cancel    context.CancelFunc
	container container.Container
}

func NewProvider(ctx context.Context, container container.Container) provider2.Provider {
	return &provider{
		ctx:       ctx,
		container: container,
	}
}

func (p provider) Name() string {
	return ProviderName
}

func (p *provider) Prepare() error {
	fmt.Println("root provider register")

	cmd := NewCmd()
	if err := cmd.Init(); err != nil {
		return err
	}
	p.container.Put(ContainerCmdName, cmd)
	return nil
}

func (p *provider) Run() error {
	fmt.Println("root provider bootstrap")

	v, ok := p.container.Get(ContainerCmdName)
	if ok {
		return v.(Command).Run()
	}
	return nil
}

func (p *provider) Shutdown() error {
	fmt.Println("root provider Shutdown")

	return nil
}
