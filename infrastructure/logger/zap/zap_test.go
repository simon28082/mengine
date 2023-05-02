package zap

import (
	"github.com/simon28082/mengine/infrastructure/logger"
	"testing"
)

func TestNewZap(t *testing.T) {
	l := NewZap(logger.Config{
		Level:   logger.DebugLevel,
		Outputs: []string{`stdout`},
	})
	//assert.Nil(t, err)
	l.Log(logger.InfoLevel, "hello", map[string]any{
		"key": "value",
		"int": 10,
		"mixed": map[string]interface{}{
			"subKey":  "subValue",
			"subInt":  10,
			"subBool": false,
		},
	})

	l.Log(logger.DebugLevel, "without params", nil)
}
func TestNewZapLevel(t *testing.T) {
	l := NewZap(logger.Config{
		Level:   logger.WarnLevel,
		Outputs: []string{`stdout`},
	})
	//assert.Nil(t, err)
	l.Log(logger.InfoLevel, "hello", map[string]any{
		"key": "value",
		"int": 10,
		"mixed": map[string]interface{}{
			"subKey":  "subValue",
			"subInt":  10,
			"subBool": false,
		},
	})

	l.Log(logger.DebugLevel, "without params", nil)

	l.Logf(logger.WarnLevel, "warn%s message", `warn`, nil)
	l.Log(logger.ErrorLevel, "err message", nil)
}
