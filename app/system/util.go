package system

import (
	"archive/tar"
	"bytes"
	"encoding/json"
	"io"

	"github.com/labstack/gommon/log"
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

func ToTar(fstream io.Reader, name string) (bytes.Buffer, error) {

	var buf bytes.Buffer
	tarW := tar.NewWriter(&buf)

	//read  all data
	data, err := io.ReadAll(fstream)

	// create and write header
	if err := tarW.WriteHeader(&tar.Header{
		Name: name,
		Size: int64(len(data)),
		Mode: 0600,
	}); err != nil {
		return bytes.Buffer{}, err
	}

	if err != nil {
		return bytes.Buffer{}, err
	}

	if _, err := tarW.Write(data); err != nil {
		log.Errorf("could not write body: %v", err)
		return bytes.Buffer{}, err
	}

	return buf, nil
}
