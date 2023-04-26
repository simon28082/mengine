package engine

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/google/wire"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
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
)

func init() {
	viper.SetEnvPrefix("MENGINE")
	viper.AutomaticEnv()
}

type Engine interface {
	Container

	Context() context.Context

	Mount(process ...Process)

	Run(args ...string) error
}

type processHandleFunc func(process Process) error

type engine struct {
	ctx       context.Context
	cancel    context.CancelFunc
	container Container
	processes []Process
	cli       *cobra.Command
}

var (
	WireEngineSet = wire.NewSet(NewEngine)

	processFunc = func(processes []Process, fn processHandleFunc, use string) error {
		for i := range processes {
			if processes[i].Global() == false {
				cobra, ok := processes[i].(CobraCommand)
				if !ok {
					continue
				}

				if cobra.Cobra().Use != use {
					continue
				}
			}
			if err := fn(processes[i]); err != nil {
				return err
			}
		}
		return nil
	}
)

func NewEngine(
	ctx context.Context,
) Engine {
	ctx, cancel := context.WithCancel(ctx)
	return &engine{
		ctx:       ctx,
		cancel:    cancel,
		container: NewContainer(),
	}
}

func (e *engine) Context() context.Context {
	return e.ctx
}

func (e *engine) Get(key string) any {
	return e.container.Get(key)
}

func (e *engine) Put(key string, val any) {
	e.container.Put(key, val)
}

func (e *engine) Add(key string, val any) bool {
	return e.container.Add(key, val)
}

func (e *engine) Delete(key string) {
	e.container.Delete(key)
}

func (e *engine) Exists(key string) bool {
	return e.container.Exists(key)
}

func (e *engine) Clean() {
	e.container.Clean()
}

func (e *engine) Mount(process ...Process) {
	e.processes = append(e.processes, process...)
}

func (e *engine) prepare() error {
	err := e.cobraBuild()
	if err != nil {
		return err
	}

	e.cobraCommandRegister()

	return nil
}

func (e *engine) Run(args ...string) error {
	if err := e.prepare(); err != nil {
		return err
	}
	//if len(args) == 0 {
	//	args = os.Args[1:]
	//}
	//e.cli.SetArgs(args)
	return e.cli.Execute()
}

func (e *engine) cobraBuild() error {
	logo1 := strings.Replace(logo, `#`, "`", 1)
	logoColor := color.New(color.FgCyan, color.Bold)
	versionColor := color.New(color.FgRed, color.Bold)
	cmd := &cobra.Command{
		Use:                `mengine`,
		Short:              `Mengine [` + version + `]`,
		Long:               fmt.Sprintf("%s Mengine [ %s ]", logoColor.Sprint(logo1), versionColor.Sprint(version)),
		Version:            version,
		PersistentPreRunE:  e.cobraPersistentPreRunE,
		PersistentPostRunE: e.cobraPersistentPostRunE,
	}

	cmd.PersistentFlags().String("log-level", viper.GetString(`LOG_LEVEL`), "log level")
	cmd.PersistentFlags().String("log-path", viper.GetString(`LOG_PATH`), "log path")
	cmd.PersistentFlags().String("config-path", "./1.json", "config level")
	e.cli = cmd
	return nil
}

func (e *engine) defaultEnvToFlags(cmd *cobra.Command) {
	var (
		logLevel = viper.GetString(`LOG_LEVEL`)
		logPath  = viper.GetString(`LOG_PATH`)
	)
	if len(logLevel) == 0 {
		logLevel = `info`
	}
	if len(logPath) == 0 {
		logPath = `stdout`
	}
	cmd.PersistentFlags().String("log-level", logLevel, "log level contains ")
	cmd.PersistentFlags().String("log-path", logPath, "log path")
	cmd.PersistentFlags().String("config-path", "./1.json", "config level")
}

func (e *engine) cobraCommandRegister() {
	for i := range e.processes {
		cmd := e.processes[i]
		if v, ok := cmd.(CobraCommand); ok {
			e.cli.AddCommand(v.Cobra())
		}
	}
}

func (e *engine) cobraPersistentPreRunE(cmd *cobra.Command, args []string) error {
	e.container.Put(`cmd`, cmd)

	var (
		logPath    = cmd.Flag("log-path").Value.String()
		logLevel   = cmd.Flag("log-level").Value.String()
		configPath = cmd.Flag(`config-path`).Value.String()
	)
	fmt.Println("logPath", logPath, "logLevel", logLevel, "configPath", configPath, "cmd.Use", cmd.Use)
	//config, err := config.NewConfig(source.NewFile(configPath))
	//if err != nil {
	//	return err
	//}
	config := configPath
	//logger := logger2.NewLogger(zap.NewZapDevelopment())

	e.container.Put(`config`, config)
	logger := ProvideZapProdLogger()
	e.container.Put(`logger`, logger)

	return processFunc(e.processes, func(process Process) error {
		return process.Prepare(e)
	}, cmd.Use)
}

func (e *engine) cobraPersistentPostRunE(cmd *cobra.Command, args []string) error {
	return processFunc(e.processes, func(process Process) error {
		return process.Shutdown(e)
	}, cmd.Use)
}
