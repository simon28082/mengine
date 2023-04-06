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

type LoggerWrap struct {
	logger Logger
}

func NewLogger(logger Logger) *LoggerWrap {
	return &LoggerWrap{
		logger: logger,
	}
}

func (l *LoggerWrap) Debug(message string, v ...interface{}) {
	l.logger.Log(Debug, message, convertToMap(v))
}

func (l *LoggerWrap) Info(message string, v ...interface{}) {
	l.logger.Log(Info, message, convertToMap(v))
}

func (l *LoggerWrap) Warn(message string, v ...interface{}) {
	l.logger.Log(Warn, message, convertToMap(v))
}

func (l *LoggerWrap) Error(message string, v ...interface{}) {
	l.logger.Log(Error, message, convertToMap(v))
}

func (l *LoggerWrap) Fatal(message string, v ...interface{}) {
	l.logger.Log(Fatal, message, convertToMap(v))
}

func (l *LoggerWrap) Panic(message string, v ...interface{}) {
	l.logger.Log(Panic, message, convertToMap(v))
}

func (l *LoggerWrap) Debugf(format, message string, v ...interface{}) {
	l.logger.Logf(Debug, format, message, convertToMap(v))
}

func (l *LoggerWrap) Infof(format, message string, v ...interface{}) {
	l.logger.Logf(Info, format, message, convertToMap(v))
}

func (l *LoggerWrap) Warnf(format, message string, v ...interface{}) {
	l.logger.Logf(Warn, format, message, convertToMap(v))
}

func (l *LoggerWrap) Errorf(format, message string, v ...interface{}) {
	l.logger.Logf(Error, format, message, convertToMap(v))
}

func (l *LoggerWrap) Panicf(format, message string, v ...interface{}) {
	l.logger.Logf(Panic, format, message, convertToMap(v))
}

func (l *LoggerWrap) Fatalf(format, message string, v ...interface{}) {
	l.logger.Logf(Fatal, format, message, convertToMap(v))
}

func (l *LoggerWrap) Log(level Level, message string, context map[string]any) {
	l.logger.Log(level, message, context)
}

func (l *LoggerWrap) Logf(level Level, format string, message string, context map[string]any) {
	l.logger.Logf(level, format, message, context)
}

func (l *LoggerWrap) Fields(m map[string]any) Logger {
	l.logger.Fields(m)
	return l
}

func (l *LoggerWrap) String() string {
	return l.logger.String()
}

func convertToMap(lists []any) map[string]any {
	listLength := len(lists)
	if listLength == 0 {
		return nil
	}

	if listLength%2 != 0 {
		lists = append(lists, nil)
	}

	var m = make(map[string]any, listLength/2)

	for i := 0; i < len(lists); i += 2 {
		m[lists[i].(string)] = lists[i+1]
	}

	return m
}
