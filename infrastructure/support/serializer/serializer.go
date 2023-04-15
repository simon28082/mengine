package serializer

import "encoding/json"

func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func Unmarshal(bytes []byte, v any) error {
	return json.Unmarshal(bytes, v)
}
