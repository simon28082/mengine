package slices

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertToMap(t *testing.T) {
	r, err := ConvertToMap(nil)
	assert.Nil(t, r)
	assert.Nil(t, err)

	r2, err1 := ConvertToMap([]any{"k1", "v1", 1, 2})
	assert.Nil(t, err1)
	assert.Equal(t, map[any]any{
		"k1": "v1",
		1:    2,
	}, r2)
	spew.Dump(r2)
	_, err3 := ConvertToMap([]any{"k1", "v1", 1})
	assert.Error(t, err3)

	r4, err1 := ConvertToMap([]any{"k1", "", 1, 2})
	assert.Nil(t, err1)
	assert.Equal(t, map[any]any{
		"k1": "",
		1:    2,
	}, r4)
}
