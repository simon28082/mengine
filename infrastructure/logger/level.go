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
		return `Debug`
	case Info:
		return `Info`
	case Warn:
		return `Warn`
	case Error:
		return `Error`
	case Panic:
		return `Panic`
	case Fatal:
		return `Fatal`
	default:
		panic(`not allow level`)
	}
}
