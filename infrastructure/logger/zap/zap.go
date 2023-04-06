package zap

import (
	"fmt"
	"github.com/simon/mengine/infrastructure/logger"
	zap2 "go.uber.org/zap"
	"sync"
)

type zap struct {
	logger *zap2.Logger
	fields sync.Map
}

func NewZap() (*zap, error) {
	zapLogger, err := zap2.NewDevelopment()
	if err != nil {
		return nil, err
	}
	return &zap{
		logger: zapLogger,
	}, nil
}

func (z *zap) Log(level logger.Level, message string, context map[string]any) {
	z.fields.Range(func(key, value any) bool {
		context[key.(string)] = value
		return true
	})

	var (
		contextLength = len(context)
		fields        []zap2.Field
	)
	if contextLength > 0 {
		fields = make([]zap2.Field, contextLength)
		var i = 0
		for key, value := range context {
			fields[i] = zap2.Any(key, value)
			i++
		}
	}

	z.logger.Log(levels.level(level), message, fields...)
}

func (z *zap) Logf(level logger.Level, format string, message string, context map[string]any) {
	z.Log(level, fmt.Sprintf(format, message), context)
}

func (z *zap) Fields(m map[string]any) logger.Logger {
	for k, v := range m {
		z.fields.Store(k, v)
	}
	return z
}

func (z *zap) String() string {
	return `zap`
}
