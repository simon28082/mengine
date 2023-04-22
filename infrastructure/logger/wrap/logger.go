package wrap

import (
	"github.com/simon28082/mengine/infrastructure/logger"
)

type LoggerWrap struct {
	logger logger.Logger
}

func NewLogger(logger logger.Logger) logger.Wrap {
	return &LoggerWrap{
		logger: logger,
	}
}

func (l *LoggerWrap) Debug(message string, v ...interface{}) {
	l.logger.Log(logger.Debug, message, convertToMap(v))
}

func (l *LoggerWrap) Info(message string, v ...interface{}) {
	l.logger.Log(logger.Info, message, convertToMap(v))
}

func (l *LoggerWrap) Warn(message string, v ...interface{}) {
	l.logger.Log(logger.Warn, message, convertToMap(v))
}

func (l *LoggerWrap) Error(message string, v ...interface{}) {
	l.logger.Log(logger.Error, message, convertToMap(v))
}

func (l *LoggerWrap) Fatal(message string, v ...interface{}) {
	l.logger.Log(logger.Fatal, message, convertToMap(v))
}

func (l *LoggerWrap) Panic(message string, v ...interface{}) {
	l.logger.Log(logger.Panic, message, convertToMap(v))
}

func (l *LoggerWrap) Debugf(format, message string, v ...interface{}) {
	l.logger.Logf(logger.Debug, format, message, convertToMap(v))
}

func (l *LoggerWrap) Infof(format, message string, v ...interface{}) {
	l.logger.Logf(logger.Info, format, message, convertToMap(v))
}

func (l *LoggerWrap) Warnf(format, message string, v ...interface{}) {
	l.logger.Logf(logger.Warn, format, message, convertToMap(v))
}

func (l *LoggerWrap) Errorf(format, message string, v ...interface{}) {
	l.logger.Logf(logger.Error, format, message, convertToMap(v))
}

func (l *LoggerWrap) Panicf(format, message string, v ...interface{}) {
	l.logger.Logf(logger.Panic, format, message, convertToMap(v))
}

func (l *LoggerWrap) Fatalf(format, message string, v ...interface{}) {
	l.logger.Logf(logger.Fatal, format, message, convertToMap(v))
}

func (l *LoggerWrap) Log(level logger.Level, message string, context map[string]any) {
	l.logger.Log(level, message, context)
}

func (l *LoggerWrap) Logf(level logger.Level, format string, message string, context map[string]any) {
	l.logger.Logf(level, format, message, context)
}

func (l *LoggerWrap) Fields(m map[string]any) logger.Logger {
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
