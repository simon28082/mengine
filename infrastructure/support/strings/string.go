package strings

import "unsafe"

func ToBytes(s string) []byte {
	length := len(s)
	if length == 0 {
		return nil
	}
	return unsafe.Slice(unsafe.StringData(s), length)
}
