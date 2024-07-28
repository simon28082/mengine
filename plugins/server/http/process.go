package http

import (
	"github.com/simon28082/mengine/infrastructure/engine"
	"github.com/urfave/cli/v3"
)

type process struct {
	engine.UnimplementedDependencies
}

func NewProcess() engine.Processor {
	return &process{}
}

func (p *process) Name() engine.ProcessorName {
	return `http.serve`
}

func (p *process) Prepare(cli *cli.Command, e *engine.Engine) (err error) {
	return nil
}

func (p *process) Start(cli *cli.Command, e *engine.Engine) (err error) {
	return nil
}

func (p *process) Shutdown(cli *cli.Command, e *engine.Engine) (err error) {
	return nil
}
