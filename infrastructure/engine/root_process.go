package engine

import (
	"context"
	"fmt"
	"github.com/google/wire"
)

const (
	ProviderName     = `root`
	ContainerCmdName = `root.engine`
)

var WireProcessSet = wire.NewSet(NewProcess, WireCmdSet)

type process struct {
	ctx    context.Context
	cancel context.CancelFunc
	cmd    Command
}

func NewProcess(ctx context.Context, cmd Command) Process {
	return &process{
		ctx: ctx,
		cmd: cmd,
	}
}

func (p *process) Name() string {
	return ProviderName
}

func (p *process) Prepare() error {
	fmt.Println("root provider register")

	p.cmd.Init()
	//cmd := NewCmd()
	//if err := cmd.Init(); err != nil {
	//	return err
	//}
	//p.container.Put(ContainerCmdName, cmd)
	return nil
}

func (p *process) Run() error {
	fmt.Println("root provider bootstrap")

	p.cmd.Run()
	//v, ok := p.container.Get(ContainerCmdName)
	//if ok {
	//	return v.(Command).Run()
	//}
	return nil
}

func (p *process) Shutdown() error {
	fmt.Println("root provider Shutdown")

	return nil
}
