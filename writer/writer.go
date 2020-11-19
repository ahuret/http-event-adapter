package writer

import (
	"fmt"

	"../configuration"
	"../log"
)

type WriterConfiguration struct {
	Logger log.Logger
	Config configuration.Provider
}

type Writer interface {
	Write(channel string, data []byte) error
}

func GetWriter(cfg *WriterConfiguration, name string) (Writer, error) {
	var writer Writer
	var err error
	switch name {
	case "fmt":
		writer, err = NewFmtWriter(cfg)
	default:
		err = fmt.Errorf("unknown writer name %s", name)
	}
	return writer, err
}