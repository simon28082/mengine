package bytes

import "unsafe"

func ToString(b []byte) string {
	length := len(b)
	if length == 0 {
		return ``
	}
	return unsafe.String(&b[0], length)
}
