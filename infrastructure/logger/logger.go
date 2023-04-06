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

	Fields(map[string]any) Logger

	String() string
}
