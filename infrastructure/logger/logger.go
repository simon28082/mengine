package logger

import (
	"io"
)

type Writer interface {
	Writer() io.Writer
}

type Logger interface {
	Log(level Level, message string, context map[string]any)

	Logf(level Level, format string, message string, context map[string]any)

	SetLevel(level Level)

	Level() Level

	String() string
}

type Wrap interface {
	Logger

	Debug(message string, v ...interface{})

	Info(message string, v ...interface{})

	Warn(message string, v ...interface{})

	Error(message string, v ...interface{})

	Fatal(message string, v ...interface{})

	Panic(message string, v ...interface{})

	Debugf(format, message string, v ...interface{})

	Infof(format, message string, v ...interface{})

	Warnf(format, message string, v ...interface{})

	Errorf(format, message string, v ...interface{})

	Panicf(format, message string, v ...interface{})

	Fatalf(format, message string, v ...interface{})
}
