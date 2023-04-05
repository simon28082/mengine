package zap

import (
	"github.com/simon/mengine/infrastructure/logger"
	"go.uber.org/zap/zapcore"
)

type levelMap map[logger.Level]zapcore.Level

var levels = levelMap{
	logger.Debug: zapcore.DebugLevel,
	logger.Info:  zapcore.InfoLevel,
	logger.Warn:  zapcore.WarnLevel,
	logger.Error: zapcore.ErrorLevel,
	logger.Panic: zapcore.PanicLevel,
	logger.Fatal: zapcore.FatalLevel,
}

func (l levelMap) level(level logger.Level) zapcore.Level {
	if v, ok := l[level]; ok {
		return v
	}
	return zapcore.DebugLevel
}
