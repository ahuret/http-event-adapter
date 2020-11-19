package format

import (
	"bytes"

	"github.com/gocarina/gocsv"
)

type csvFormatter struct{}

func NewCsvFormatter(cfg *FormatterConfiguration) (Formatter, error) {
	return &csvFormatter{}, nil
}

func (c *csvFormatter) Format(data []byte) (interface{}, error) {
	r := bytes.NewReader(data)
	return gocsv.CSVToMaps(r)
}
