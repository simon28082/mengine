package logger

type WrapLogger struct {
	logger Logger
}

func NewLogger(l Logger) Wrap {
	return &WrapLogger{
		logger: l,
	}
}

func (l *WrapLogger) Debug(message string, v ...interface{}) {
	l.logger.Log(DebugLevel, message, convertToMap(v))
}

func (l *WrapLogger) Info(message string, v ...interface{}) {
	l.logger.Log(InfoLevel, message, convertToMap(v))
}

func (l *WrapLogger) Warn(message string, v ...interface{}) {
	l.logger.Log(WarnLevel, message, convertToMap(v))
}

func (l *WrapLogger) Error(message string, v ...interface{}) {
	l.logger.Log(ErrorLevel, message, convertToMap(v))
}

func (l *WrapLogger) Fatal(message string, v ...interface{}) {
	l.logger.Log(FatalLevel, message, convertToMap(v))
}

func (l *WrapLogger) Panic(message string, v ...interface{}) {
	l.logger.Log(PanicLevel, message, convertToMap(v))
}

func (l *WrapLogger) Debugf(format, message string, v ...interface{}) {
	l.logger.Logf(DebugLevel, format, message, convertToMap(v))
}

func (l *WrapLogger) Infof(format, message string, v ...interface{}) {
	l.logger.Logf(InfoLevel, format, message, convertToMap(v))
}

func (l *WrapLogger) Warnf(format, message string, v ...interface{}) {
	l.logger.Logf(WarnLevel, format, message, convertToMap(v))
}

func (l *WrapLogger) Errorf(format, message string, v ...interface{}) {
	l.logger.Logf(ErrorLevel, format, message, convertToMap(v))
}

func (l *WrapLogger) Panicf(format, message string, v ...interface{}) {
	l.logger.Logf(PanicLevel, format, message, convertToMap(v))
}

func (l *WrapLogger) Fatalf(format, message string, v ...interface{}) {
	l.logger.Logf(FatalLevel, format, message, convertToMap(v))
}

func (l *WrapLogger) Log(level Level, message string, context map[string]any) {
	l.logger.Log(level, message, context)
}

func (l *WrapLogger) Logf(level Level, format string, message string, context map[string]any) {
	l.logger.Logf(level, format, message, context)
}

func (l *WrapLogger) SetLevel(level Level) {
	l.logger.SetLevel(level)
}

func (l *WrapLogger) Level() Level {
	return l.logger.Level()
}

func (l *WrapLogger) String() string {
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
