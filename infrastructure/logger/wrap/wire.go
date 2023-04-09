package wrap

import (
	"github.com/google/wire"
	"github.com/simon/mengine/infrastructure/logger/zap"
	"sync"
)

var (
	WireZapLoggerDevelopmentSet = wire.NewSet(ZapLoggerDevelopmentWrapOnce, wire.Value(ZapLoggerDevelopmentWrapOnce))
	zapLoggerDevelopmentOnce    sync.Once
	zapLoggerDevelopment        *LoggerWrap
)

//func ZapLoggerDevelopmentWrapFunc() *LoggerWrap {
//	//panic(wire.Build(NewLogger, zap.WireZapDevelopmentSet))
//}

func ZapLoggerDevelopmentWrapOnce() *LoggerWrap {
	if zapLoggerDevelopment == nil {
		zapLoggerDevelopmentOnce.Do(func() {
			zapLoggerDevelopment = NewLogger(zap.NewZapDevelopment())
		})
	}

	return zapLoggerDevelopment
}

var (
	WireZapLoggerProductionSet = wire.NewSet(ZapLoggerProductionWrapOnce, wire.Value(ZapLoggerProductionWrapOnce))
	zapLoggerProductionOnce    sync.Once
	zapLoggerProduction        *LoggerWrap
)

//func ZapLoggerProductionWrapFunc() *LoggerWrap {
//	panic(wire.Build(NewLogger, zap.WireZapProductionSet))
//}

func ZapLoggerProductionWrapOnce() *LoggerWrap {
	if zapLoggerProduction == nil {
		zapLoggerProductionOnce.Do(func() {
			zapLoggerProduction = NewLogger(zap.NewZapProduction())
		})
	}

	return zapLoggerProduction
}
