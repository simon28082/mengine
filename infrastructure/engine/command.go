package engine

import (
	"context"
	"fmt"
	"github.com/simon28082/mengine/infrastructure/support/option"
	"github.com/urfave/cli/v3"
)

func init() {
	cli.VersionPrinter = func(cmd *cli.Command) {
		fmt.Printf("%s:%s", cmd.Name, cmd.Version)
	}
}

type Command struct {
	options    CommandOptions
	cmd        *cli.Command
	processors []Processor
}

type CommandOptions struct {
	Context      context.Context
	Cancel       context.CancelFunc
	Name         string
	Usage        string
	UsageText    string
	Version      string
	Action       cli.ActionFunc
	BeforeAction cli.BeforeFunc
	AfterAction  cli.AfterFunc
	Flags        []cli.Flag
}

func WithContext(ctx context.Context) option.Option[*CommandOptions] {
	return func(opts *CommandOptions) {
		opts.Context, opts.Cancel = context.WithCancel(ctx)
	}
}

func WithAction(fn cli.ActionFunc) option.Option[*CommandOptions] {
	return func(opts *CommandOptions) {
		opts.Action = fn
	}
}

func WithBeforeAction(fn cli.BeforeFunc) option.Option[*CommandOptions] {
	return func(opts *CommandOptions) {
		opts.BeforeAction = fn
	}
}

func WithAfterAction(fn cli.AfterFunc) option.Option[*CommandOptions] {
	return func(opts *CommandOptions) {
		opts.AfterAction = fn
	}
}

func WithFlag(flags ...cli.Flag) option.Option[*CommandOptions] {
	return func(opts *CommandOptions) {
		opts.Flags = append(opts.Flags, flags...)
	}
}

func WithName(name string) option.Option[*CommandOptions] {
	return func(opts *CommandOptions) {
		opts.Name = name
	}
}

func WithUsage(usage string) option.Option[*CommandOptions] {
	return func(opts *CommandOptions) {
		opts.Usage = usage
	}
}

func WithUsageText(text string) option.Option[*CommandOptions] {
	return func(opts *CommandOptions) {
		opts.UsageText = text
	}
}

func WithVersion(version string) option.Option[*CommandOptions] {
	return func(opts *CommandOptions) {
		opts.Version = version
	}
}

func NewCommand(opts ...option.Option[*CommandOptions]) *Command {
	var cmdOptions CommandOptions
	option.Apply(&cmdOptions, opts...)

	if cmdOptions.Context == nil {
		cmdOptions.Context = context.Background()
	}

	cmd := &cli.Command{
		Name:      cmdOptions.Name,
		Flags:     cmdOptions.Flags,
		Action:    cmdOptions.Action,
		Before:    cmdOptions.BeforeAction,
		After:     cmdOptions.AfterAction,
		Usage:     cmdOptions.Usage,
		UsageText: cmdOptions.UsageText,
	}

	return &Command{
		cmd:        cmd,
		options:    cmdOptions,
		processors: make([]Processor, 0),
	}
}

func (c *Command) Options() CommandOptions {
	return c.options
}

func (c *Command) Cmd() *cli.Command {
	return c.cmd
}

func (c *Command) Processors() []Processor {
	var nps = make([]Processor, len(c.processors))
	copy(nps, c.processors)
	return nps
}

func (c *Command) Mount(ps ...Processor) error {
	c.processors = append(c.processors, ps...)
	return nil
}
