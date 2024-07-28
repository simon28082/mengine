package engine

import (
	"context"
	"github.com/simon28082/mengine/infrastructure/errors"
	"github.com/simon28082/mengine/infrastructure/support/option"
	slices2 "github.com/simon28082/mengine/infrastructure/support/slices"
	"github.com/urfave/cli/v3"
	"go.uber.org/dig"
	"log/slog"
	"os"
	"sync"
)

const (
	logo = `
 __  __                  _
|  \/  | ___ _ __   __ _(_)_ __   ___
| |\/| |/ _ \ '_ \ / _# | | '_ \ / _ \
| |  | |  __/ | | | (_| | | | | |  __/
|_|  |_|\___|_| |_|\__, |_|_| |_|\___|
                   |___/
	
`

	version = `v0.0.1-alpha`
	name    = `mengine`
)

type commandProcessor struct {
	processors []Processor
	command    *Command
}

type Engine struct {
	root        *cli.Command
	options     CommandOptions
	processors  []Processor
	commands    map[string]*commandProcessor
	processLock sync.RWMutex
	dig         *dig.Container
}

//type Options struct {
//	Name    string
//	Version string
//	Context context.Context
//	Cancel  context.CancelFunc
//}

//func Name(name string) option.Option[*Options] {
//	return func(opts *Options) {
//		opts.Name = name
//	}
//}
//
//func Context(ctx context.Context) option.Option[*Options] {
//	return func(opts *Options) {
//		opts.Context, opts.Cancel = context.WithCancel(ctx)
//	}
//}
//
//func Version(version string) option.Option[*Options] {
//	return func(opts *Options) {
//		opts.Version = version
//	}
//}

func NewEngine(opts ...option.Option[*CommandOptions]) *Engine {
	var newOptions = []option.Option[*CommandOptions]{
		WithName(name),
		WithVersion(version),
		WithUsageText(logo),
	}

	newOptions = append(newOptions, opts...)

	command := NewCommand(newOptions...)

	return &Engine{
		root:       command.Cmd(),
		dig:        dig.New(),
		options:    command.options,
		processors: make([]Processor, 0),
		commands:   make(map[string]*commandProcessor),
	}
}

func (e *Engine) Context() context.Context {
	return e.options.Context
}

func (e *Engine) Provide(constructs ...interface{}) error {
	for i := range constructs {
		var (
			construct = constructs[i]
			err       = e.dig.Provide(construct)
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *Engine) Decorate(constructs ...interface{}) error {
	for i := range constructs {
		var (
			construct = constructs[i]
			err       = e.dig.Decorate(construct)
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *Engine) Invoke(fns ...interface{}) error {
	for i := range fns {
		var (
			fn  = fns[i]
			err = e.dig.Invoke(fn)
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *Engine) Mount(ps ...Processor) error {
	e.processors = append(e.processors, ps...)
	return nil
}

func (e *Engine) AddCommand(command *Command) error {
	cmdName := command.Options().Name
	if _, ok := e.commands[cmdName]; ok {
		return errors.Errorf(`command:%s already exists`, cmdName)
	}

	cpr := &commandProcessor{
		command:    command,
		processors: make([]Processor, 0),
	}

	if err := e.buildCommand(cpr); err != nil {
		return err
	}
	e.commands[cmdName] = cpr
	return nil
}

func (e *Engine) Run(args ...string) error {
	for _, v := range e.commands {
		e.root.Commands = append(e.root.Commands, v.command.Cmd())
	}

	if len(args) == 0 {
		args = os.Args
	}

	return e.root.Run(e.Context(), args)
}

func (e *Engine) Close() error {
	e.options.Cancel()
	return nil
}

func (e *Engine) allCommandProcessors(command *Command) []Processor {
	return append(e.processors, command.Processors()...)
}

func (e *Engine) resolveCommandProcessors(command *Command) ([]Processor, error) {
	return e.resolveDependency(e.allCommandProcessors(command))
}

func (e *Engine) buildCommand(cpr *commandProcessor) error {
	var (
		cmd              = cpr.command.Cmd()
		beforeActionFunc = cmd.Before
		actionFunc       = cmd.Action
		afterActionFunc  = cmd.After
	)

	cmd.Before = func(ctx context.Context, command *cli.Command) error {
		if beforeActionFunc != nil {
			return beforeActionFunc(ctx, command)
		}

		return nil
	}
	cmd.Action = func(ctx context.Context, command *cli.Command) (err error) {
		cpr.processors, err = e.resolveCommandProcessors(cpr.command)
		if err != nil {
			return err
		}

		err = e.prepareProcessor(command, cpr.processors)
		if err != nil {
			return err
		}

		err = e.startProcessor(command, cpr.processors)
		if err != nil {
			return err
		}

		if actionFunc != nil {
			errn := actionFunc(ctx, command)
			if errn != nil {
				return errn
			}
		}

		return nil
	}
	cmd.After = func(ctx context.Context, command *cli.Command) error {
		if afterActionFunc != nil {
			if err := afterActionFunc(ctx, command); err != nil {
				slog.Error(`shutdown command failed`, `error`, err)
			}
		}

		if err := e.shutdownProcessor(command, cpr.processors); err != nil {
			slog.Error(`process shutdown failed`, `error`, err)
		}

		return nil
	}

	return nil
}

func (e *Engine) prepareProcessor(command *cli.Command, processors []Processor) error {
	for _, p := range processors {
		if err := p.Prepare(command, e); err != nil {
			return errors.Errorf(err, `processor:%s prepare failed`, p.Name())
		} else {
			slog.Info(`prepare processor`, `name`, p.Name())
		}
	}

	return nil
}

func (e *Engine) startProcessor(command *cli.Command, processors []Processor) error {
	for _, p := range processors {
		if err := p.Start(command, e); err != nil {
			return errors.Errorf(err, `processor:%s start failed`, p.Name())
		} else {
			slog.Info(`start processor`, `name`, p.Name())
		}
	}

	return nil
}

func (e *Engine) shutdownProcessor(command *cli.Command, processors []Processor) error {
	for _, p := range slices2.Reverse(processors) {
		if err := p.Shutdown(command, e); err != nil {
			return errors.Errorf(err, `processor:%s shutdown failed`, p.Name())
		} else {
			slog.Info(`shutdown processor`, `name`, p.Name())
		}
	}

	return nil
}

func (e *Engine) resolveDependency(ps []Processor) ([]Processor, error) {
	var (
		walk              func(p Processor) error
		processorLength   = len(ps)
		foundedProcessors = make(map[ProcessorName]struct{}, processorLength)
		allProcessorNames = make(map[ProcessorName]Processor, processorLength)
		orderProcessors   []Processor
	)

	for _, p := range ps {
		allProcessorNames[p.Name()] = p
	}

	walk = func(p Processor) error {
		// skip processors that have been processed
		if _, found := foundedProcessors[p.Name()]; found {
			return nil
		}

		foundedProcessors[p.Name()] = struct{}{}

		for _, name := range p.Dependencies() {
			if v, ok := allProcessorNames[name]; !ok {
				return errors.Errorf(`the processor:%s depends on a process that doesn't exist`, name)
			} else {
				if err := walk(v); err != nil {
					return err
				}
			}
		}

		orderProcessors = append(orderProcessors, p)

		return nil
	}

	for _, p := range ps {
		if err := walk(p); err != nil {
			return nil, err
		}
	}

	return orderProcessors, nil
}
