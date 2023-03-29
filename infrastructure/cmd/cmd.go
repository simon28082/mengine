package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

type Command interface {
	Run() error

	AddCommand(Command) error
}

type Cmd struct {
	cli *cobra.Command
}

func NewCmd() *Cmd {
	return &Cmd{
		cli: mengineCommand(),
	}
}

func (c *Cmd) AddCommand() {

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

	//cmd.PersistentFlags().StringP("config", "c", getConfigPath(configPath), "config path")
	//cmd.PersistentFlags().BoolP("dev", "", false, "open development mode (default production)")

	//cmd.SetVersionTemplate("{{with .Short}}{{printf \"%s \" .}}{{end}}{{printf \"Version %s\" .Version}}\n")

	return cmd
}

func (c *Cmd) Run() error {
	c.cli.SetArgs(os.Args[1:])
	c.cli.AddCommand(&cobra.Command{
		Use: `test`,
	})
	return c.cli.Execute()
}
