package logger

type Level int8

const (
	Debug Level = 0
	Info  Level = 1
	Warn  Level = 2
	Error Level = 3
	Panic Level = 4
	Fatal Level = 5
)

func (l Level) String() string {
	switch l {
	case Debug:
		return `debug`
	case Info:
		return `info`
	case Warn:
		return `warn`
	case Error:
		return `error`
	case Panic:
		return `panic`
	case Fatal:
		return `fatal`
	default:
		panic(`not allow level`)
	}
}
