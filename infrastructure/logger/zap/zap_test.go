package zap

import (
	"github.com/simon/mengine/infrastructure/logger"
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
}
