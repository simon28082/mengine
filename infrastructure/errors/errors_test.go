package errors

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewDefault(t *testing.T) {
	err1 := NewDefault("abc")
	err2 := WithErrorf(err1, "err2")
	err3 := WithErrorf("err3", err1, err2)
	println(err3.Error())

}

func TestErrorf(t *testing.T) {
	abc := fmt.Sprintf(``, "errorfdasfdafdasfdsa")
	fmt.Println(abc)
	os.Exit(0)
	err1 := NewDefault("abc")
	err2 := WithErrorf(err1, "format %s", "----")
	fmt.Println(err2)
	assert.Contains(t, err2.Error(), `format ----`)
	err3 := WithError(err1, "format ------", "abc")
	assert.Contains(t, err3.Error(), `format ------`)
	fmt.Println(err3.Error())
}
