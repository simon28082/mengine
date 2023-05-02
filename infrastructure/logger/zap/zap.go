package zap

import (
	"fmt"
	"github.com/simon28082/mengine/infrastructure/logger"
	zap2 "go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type zap struct {
	logger *zap2.Logger
	config *logger.Config
}

//var (
//	WireZapDevelopmentSet = wire.NewSet(NewZapDevelopment)
//	WireZapProductionSet  = wire.NewSet(NewZapProduction)
//)

//func NewZapDevelopment() *zap {
//	zapLogger, err := zap2.NewDevelopment()
//	if err != nil {
//		panic(err)
//	}
//	return &zap{
//		logger: zapLogger,
//	}
//}
//
//func NewZapProduction() *zap {
//	zapLogger, err := zap2.NewProduction()
//	if err != nil {
//		panic(err)
//	}
//	return &zap{
//		logger: zapLogger,
//	}
//}

func SetZapDefaultLogger(config logger.Config) {
	logger.DefaultLogger = logger.NewLogger(NewZap(config))
}

func NewZap(config logger.Config) *zap {
	config = logger.DefaultConfigMerge(config)

	var (
		zapLevel      = levels[config.Level]
		zapTraceLevel = levels[config.TraceLevel]
		encoderConfig = zap2.NewProductionEncoderConfig()
		zc            zapcore.Core
	)

	encoderConfig.TimeKey = `time`
	encoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	//encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	if len(config.Outputs) == 1 && logger.IsConsoleOutput(config.Outputs[0]) {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		zc = zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zap2.CombineWriteSyncers(os.Stdout), zapLevel)
	} else {
		// get first file path
		var file string
		for _, f1 := range config.Outputs {
			if logger.IsConsoleOutput(f1) {
				continue
			}
			file = f1
			break
		}
		f := &lumberjack.Logger{
			Filename:   file,
			MaxSize:    config.FileMaxSize,
			MaxBackups: config.FileMaxBackups,
			MaxAge:     config.FileMaxAge,
		}
		zc = zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(f), zapLevel)
	}
	zapLogger := zap2.New(zc,
		zap2.AddStacktrace(zapTraceLevel),
		zap2.AddCallerSkip(config.Skip),
		zap2.AddCaller(),
	)

	if len(config.Fields) > 0 {
		fields := make([]zap2.Field, len(config.Fields))
		var i = 0
		for key, value := range config.Fields {
			fields[i] = zap2.Any(key, value)
			i++
		}
		zc.With(fields)
	}

	return &zap{
		config: &config,
		logger: zapLogger,
	}
}

func (z *zap) SetLevel(level logger.Level) {
	z2 := z.logger.Level()
	if err := z2.Set(level.String()); err != nil {
		panic(err)
	}
}

func (z *zap) Level() logger.Level {
	return reverseLevel.level(z.logger.Level())
}

func (z *zap) Log(level logger.Level, message string, context map[string]any) {
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

	z.logger.Log(levels.level(level), fmt.Sprintf(format, message), fields...)
}

func (z *zap) String() string {
	return `zap`
}
