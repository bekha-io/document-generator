package types

import (
	"bytes"
	"io"
)

type TemplateFile struct {
	Data []byte
}

func (t TemplateFile) ReadWriter() io.ReadWriter {
	buf := bytes.NewBuffer(t.Data)
	return buf
}