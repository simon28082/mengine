package logger

var DefaultLogger Wrap

func Debug(message string, v ...interface{}) {
	if DefaultLogger != nil {
		DefaultLogger.Debug(message, v...)
	}
}

func Info(message string, v ...interface{}) {
	if DefaultLogger != nil {
		DefaultLogger.Info(message, v...)
	}
}

func Warn(message string, v ...interface{}) {
	if DefaultLogger != nil {
		DefaultLogger.Warn(message, v...)
	}
}

func Error(message string, v ...interface{}) {
	if DefaultLogger != nil {
		DefaultLogger.Error(message, v...)
	}
}

func Fatal(message string, v ...interface{}) {
	if DefaultLogger != nil {
		DefaultLogger.Fatal(message, v...)
	}
}

func Panic(message string, v ...interface{}) {
	if DefaultLogger != nil {
		DefaultLogger.Panic(message, v...)
	}
}

func Debugf(format string, message string, v ...interface{}) {
	if DefaultLogger != nil {
		DefaultLogger.Debugf(format, message, v...)
	}
}

func Infof(format, message string, v ...interface{}) {
	if DefaultLogger != nil {
		DefaultLogger.Infof(format, message, v...)
	}
}

func Warnf(format, message string, v ...interface{}) {
	if DefaultLogger != nil {
		DefaultLogger.Warnf(format, message, v...)
	}
}

func Errorf(format, message string, v ...interface{}) {
	if DefaultLogger != nil {
		DefaultLogger.Errorf(format, message, v...)
	}
}

func Panicf(format, message string, v ...interface{}) {
	if DefaultLogger != nil {
		DefaultLogger.Panicf(format, message, v...)
	}
}

func Fatalf(format, message string, v ...interface{}) {
	if DefaultLogger != nil {
		DefaultLogger.Fatalf(format, message, v...)
	}
}
