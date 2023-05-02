package zap

import (
	"github.com/simon28082/mengine/infrastructure/logger"
	"go.uber.org/zap/zapcore"
)

type levelMap map[logger.Level]zapcore.Level
type reverseLevelMap map[zapcore.Level]logger.Level

var levels = levelMap{
	logger.DebugLevel: zapcore.DebugLevel,
	logger.InfoLevel:  zapcore.InfoLevel,
	logger.WarnLevel:  zapcore.WarnLevel,
	logger.ErrorLevel: zapcore.ErrorLevel,
	logger.PanicLevel: zapcore.PanicLevel,
	logger.FatalLevel: zapcore.FatalLevel,
}

var reverseLevel = reverseLevelMap{
	zapcore.DebugLevel: logger.DebugLevel,
	zapcore.InfoLevel:  logger.InfoLevel,
	zapcore.WarnLevel:  logger.WarnLevel,
	zapcore.ErrorLevel: logger.ErrorLevel,
	zapcore.PanicLevel: logger.PanicLevel,
	zapcore.FatalLevel: logger.FatalLevel,
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
	return logger.DebugLevel
}
