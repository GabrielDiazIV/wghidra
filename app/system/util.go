package system

import (
	"archive/tar"
	"bytes"
	"encoding/json"
	"io"

	"github.com/hashicorp/go-multierror"
	"github.com/labstack/gommon/log"
)

func Encode[T any](t T) (io.Reader, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	if err := encoder.Encode(t); err != nil {
		return nil, err
	}
	return &b, nil
}

func Decode[T any](reader io.Reader) (T, error) {
	var msg T
	decoder := json.NewDecoder(reader)
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

type MultiCloser struct {
	closers []io.Closer
}

func (m *MultiCloser) Close() error {
	var err error
	for _, c := range m.closers {
		if e := c.Close(); e != nil {
			err = multierror.Append(err, e)
		}
	}
	return err
}

func NewMultiCloser(closers []io.Closer) *MultiCloser {
	return &MultiCloser{
		closers: closers,
	}

}

func GetReaders(srcReader io.Reader, size int) []io.Reader {
	readers := make([]io.Reader, size)
	pipeWriters := make([]io.Writer, size)
	pipeClosers := make([]io.Closer, size)

	for i := 0; i < size; i++ {
		pr, pw := io.Pipe()
		readers[i] = pr
		pipeWriters[i] = pw
		pipeClosers[i] = pw
	}

	multiWriter := io.MultiWriter(pipeWriters...)
	multiCloser := NewMultiCloser(pipeClosers)

	go func() {
		io.Copy(multiWriter, srcReader)
		multiCloser.Close()
	}()

	return readers
}
