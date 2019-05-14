package reader

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

// Reader is the main resource reader & loader for flyff's files
type Reader struct {
	Filename    string
	Bytes       []byte
	BytesReader *bytes.Reader
}

// ReadAll file & return a buffer or error
func (r *Reader) ReadAll() error {
	b, err := ioutil.ReadFile(fmt.Sprintf("resources/%s", r.Filename))
	if err != nil {
		return err
	}

	r.BytesReader = bytes.NewReader(b)
	r.Bytes = b
	return err
}
