package writer

import (
	gofmt "fmt"
)

type fmtWriter struct{}

func NewFmtWriter(cfg *WriterConfiguration) (Writer, error) {
	return &fmtWriter{}, nil
}

func (w *fmtWriter) Write(channel string, data []byte) error {
	gofmt.Printf("---\nnew event sent on %s\n%s\n---\n", channel, string(data))
	return nil
}
