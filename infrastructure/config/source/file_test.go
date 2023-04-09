package source

import (
	"fmt"
	os2 "github.com/simon/mengine/infrastructure/support/os"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"path"
	"testing"
)

func TestNewFile(t *testing.T) {
	var filepath = path.Join(os2.RunDir(), `testdata/1.json`)
	fmt.Println(filepath)
	file := NewFile(filepath)
	rd, err := file.Read()
	assert.Nil(t, err)
	fsc, err := os.ReadFile(filepath)
	rs, err1 := io.ReadAll(rd)
	assert.Nil(t, err1)
	assert.Equal(t, fsc, rs)
}
