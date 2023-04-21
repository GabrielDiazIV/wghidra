package system

import (
	"bytes"
	"encoding/json"
)

func Encode[T any](t T) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	if err := encoder.Encode(t); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func Decode[T any](data []byte) (T, error) {
	var msg T
	buf := bytes.NewBuffer(data)
	decoder := json.NewDecoder(buf)
	err := decoder.Decode(&msg)
	return msg, err
}
