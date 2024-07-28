package slices

import "errors"

func ConvertToMap(lists []any) (map[any]any, error) {
	listLength := len(lists)
	if listLength == 0 {
		return nil, nil
	}

	if listLength%2 != 0 {
		return nil, errors.New(`list length does not match`)
	}

	var m = make(map[any]any, listLength/2)

	for i := 0; i < len(lists); i += 2 {
		m[lists[i]] = lists[i+1]
	}

	return m, nil
}

func InSlice[T comparable](key T, targets []T) bool {
	for i := range targets {
		if targets[i] == key {
			return true
		}
	}
	return false
}

func Reverse[T any](s []T) []T {
	oldLen := len(s)
	if oldLen == 0 {
		return s
	}

	var s1 = make([]T, 0, len(s))
	for i := oldLen - 1; i >= 0; i-- {
		s1 = append(s1, s[i])
	}

	return s1
}
