package engine

import "github.com/urfave/cli/v3"

type ProcessorName string

type Processor interface {
	Name() ProcessorName

	Dependencies() []ProcessorName

	Prepare(cli *cli.Command, e *Engine) (err error)

	Start(cli *cli.Command, e *Engine) (err error)

	Shutdown(cli *cli.Command, e *Engine) (err error)
}

type UnimplementedDependencies struct{}

func (UnimplementedDependencies) Dependencies() []ProcessorName {
	return nil
}
