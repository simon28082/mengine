package http

import (
	"fmt"
	"github.com/google/wire"
	cmd2 "github.com/simon/mengine/infrastructure/engine"
	"github.com/spf13/cobra"
)

var WireCmdSet = wire.NewSet(wire.InterfaceValue(new(cmd2.Command), NewCmd()))

type cmd struct {
	cli *cobra.Command
}

func NewCmd() *cmd {
	return &cmd{}
}

func (c *cmd) Init() error {
	cli := &cobra.Command{
		Use: `http`,
	}
	cli.Run = func(cmd *cobra.Command, args []string) {
		fmt.Println("http start..........")
	}

	c.cli = cli

	return nil
}

func (c *cmd) Cobra() *cobra.Command {
	return c.cli
}

func (c *cmd) Run(args ...string) error {
	return nil
	//return c.cli.Execute()
}

func (c cmd) AddCommand(cmds ...cmd2.Command) error {
	//TODO implement me
	panic("implement me")
}
