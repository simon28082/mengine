package logger

type Level int8

const (
	DebugLevel Level = 0
	InfoLevel  Level = 1
	WarnLevel  Level = 2
	ErrorLevel Level = 3
	PanicLevel Level = 4
	FatalLevel Level = 5
)

func (l Level) String() string {
	switch l {
	case DebugLevel:
		return `debug`
	case InfoLevel:
		return `info`
	case WarnLevel:
		return `warn`
	case ErrorLevel:
		return `error`
	case PanicLevel:
		return `panic`
	case FatalLevel:
		return `fatal`
	default:
		panic(`not allow level`)
	}
}

func StringLevel(l string) Level {
	switch l {
	case `debug`, `DEBUG`:
		return DebugLevel
	case `info`, `INFO`:
		return InfoLevel
	case `warn`, `WARN`, `warning`, `WARNING`:
		return WarnLevel
	case `error`, `ERROR`:
		return ErrorLevel
	case `panic`, `PANIC`:
		return PanicLevel
	case `fatal`, `FATAL`:
		return FatalLevel
	default:
		panic(`not allow level string`)
	}
}
