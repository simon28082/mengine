package zap

import (
	"github.com/simon28082/mengine/infrastructure/logger"
	"go.uber.org/zap/zapcore"
)

type levelMap map[logger.Level]zapcore.Level
type reverseLevelMap map[zapcore.Level]logger.Level

var levels = levelMap{
	logger.Debug: zapcore.DebugLevel,
	logger.Info:  zapcore.InfoLevel,
	logger.Warn:  zapcore.WarnLevel,
	logger.Error: zapcore.ErrorLevel,
	logger.Panic: zapcore.PanicLevel,
	logger.Fatal: zapcore.FatalLevel,
}

var reverseLevel = reverseLevelMap{
	zapcore.DebugLevel: logger.Debug,
	zapcore.InfoLevel:  logger.Info,
	zapcore.WarnLevel:  logger.Warn,
	zapcore.ErrorLevel: logger.Error,
	zapcore.PanicLevel: logger.Panic,
	zapcore.FatalLevel: logger.Fatal,
}

func (l levelMap) level(level logger.Level) zapcore.Level {
	if v, ok := l[level]; ok {
		return v
	}
	return zapcore.DebugLevel

}

func (rl reverseLevelMap) level(level zapcore.Level) logger.Level {
	if v, ok := rl[level]; ok {
		return v
	}
	return logger.Debug
}
