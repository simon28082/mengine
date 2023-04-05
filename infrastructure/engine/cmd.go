package engine

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/google/wire"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

type Command interface {
	Init() error

	AddCommand(cmds ...Command) error

	Run(args ...string) error
}

type Cmd struct {
	cli      *cobra.Command
	commands []Command
}

var WireCmdSet = wire.NewSet(wire.InterfaceValue(new(Command), NewCmd()))

func NewCmd() *Cmd {
	return &Cmd{
		cli:      mengineCommand(),
		commands: make([]Command, 0),
	}
}

func (c *Cmd) Init() error {
	return nil
}

func (c *Cmd) AddCommand(cmds ...Command) error {
	for i := range cmds {
		cmd := cmds[i]
		if v, ok := cmd.(CobraCommand); ok {

			c.cli.AddCommand(v.Cobra())
			println("register================================================================")
		}
	}
	return nil
}

func mengineCommand() *cobra.Command {
	logo := `
 __  __                  _
|  \/  | ___ _ __   __ _(_)_ __   ___
| |\/| |/ _ \ '_ \ / _# | | '_ \ / _ \
| |  | |  __/ | | | (_| | | | | |  __/
|_|  |_|\___|_| |_|\__, |_|_| |_|\___|
                   |___/
	
`
	logo = strings.Replace(logo, `#`, "`", 1)
	logoColor := color.New(color.FgCyan, color.Bold)
	versionColor := color.New(color.FgRed, color.Bold)
	version := `1.1`
	cmd := &cobra.Command{
		Use:     `mengine`,
		Short:   `Mengine [` + version + `]`,
		Long:    fmt.Sprintf("%s Mengine [ %s ]", logoColor.Sprint(logo), versionColor.Sprint(version)),
		Version: version,
	}

	//engine.PersistentFlags().StringP("config", "c", getConfigPath(configPath), "config path")
	//engine.PersistentFlags().BoolP("dev", "", false, "open development mode (default production)")

	//engine.SetVersionTemplate("{{with .Short}}{{printf \"%s \" .}}{{end}}{{printf \"Version %s\" .Version}}\n")

	return cmd
}

func (c *Cmd) Run(args ...string) error {
	if len(args) == 0 {
		args = os.Args[1:]
	}
	c.cli.SetArgs(args)
	return c.cli.Execute()
}
