package errors

import "testing"

func TestNewDefault(t *testing.T) {
	err1 := NewDefault("abc")
	err2 := WithErrorf("err2", err1)
	err3 := WithErrorf("err3", err1, err2)
	println(err3.Error())

}
