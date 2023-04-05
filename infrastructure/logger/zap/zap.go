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
	var fields = make([]zap2.Field, len(context))
	z.fields.Range(func(key, value any) bool {
		fields = append(fields, zap2.Any(key.(string), value))
		return true
	})
	if len(context) > 0 {
		for k, v := range context {
			fields = append(fields, zap2.Any(k, v))
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
