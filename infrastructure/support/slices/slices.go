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
