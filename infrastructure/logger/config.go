package logger

var (
	consoleOutput = map[string]struct{}{
		`stdout`: {},
		`stderr`: {},
	}

	defaultSkip = 3
)

func IsConsoleOutput(output string) bool {
	_, ok := consoleOutput[output]
	return ok
}

type Config struct {
	Level      Level
	TraceLevel Level
	Skip       int
	Outputs    []string
	// FileMaxAge max save days
	FileMaxAge int
	// FileMaxBackups Maximum number of files backed up
	FileMaxBackups int
	// FileMaxSize max size units:MB
	FileMaxSize int
	Fields      map[string]any
}

func DefaultFileConfig() Config {
	var (
		maxDay     = 10
		maxSize    = 200
		maxBackups = 10
	)

	return Config{
		Level:          InfoLevel,
		TraceLevel:     WarnLevel,
		Skip:           defaultSkip,
		FileMaxSize:    maxSize,
		FileMaxBackups: maxBackups,
		FileMaxAge:     maxDay,
		Outputs:        []string{`/var/log/mengine/run.log`},
	}
}

func DefaultConfig() Config {
	return Config{
		Level:      InfoLevel,
		TraceLevel: WarnLevel,
		Skip:       defaultSkip,
		Outputs:    []string{`stdout`},
	}
}

func DefaultDebugConfig() Config {
	return Config{
		Level:      DebugLevel,
		TraceLevel: WarnLevel,
		Skip:       defaultSkip,
		Outputs:    []string{`stdout`},
	}
}

func DefaultConfigMerge(config Config) Config {
	var fileConfig = DefaultFileConfig()
	if config.FileMaxAge == 0 {
		config.FileMaxAge = fileConfig.FileMaxAge
	}
	if config.FileMaxSize == 0 {
		config.FileMaxSize = fileConfig.FileMaxSize
	}
	if config.FileMaxBackups == 0 {
		config.FileMaxBackups = fileConfig.FileMaxBackups
	}
	if len(config.Outputs) == 0 {
		config.Outputs = fileConfig.Outputs
	}
	if config.Skip == 0 {
		config.Skip = fileConfig.Skip
	}
	return config
}
