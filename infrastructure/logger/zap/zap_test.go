package zap

import (
	"github.com/simon28082/mengine/infrastructure/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewZap(t *testing.T) {
	l, err := NewZap()
	assert.Nil(t, err)
	l.Log(logger.Debug, "hello", map[string]any{
		"key": "value",
		"int": 10,
		"mixed": map[string]interface{}{
			"subKey":  "subValue",
			"subInt":  10,
			"subBool": false,
		},
	})

	l.Log(logger.Debug, "without params", nil)
}
