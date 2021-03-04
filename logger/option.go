package logger

import (
	"io"
	"os"
)

// Option for logger
type Option struct {
	AppName string
	Level   string
	Output  io.Writer
}

// DefaultOption contains default value for some attributes
func DefaultOption() *Option {
	return &Option{
		Level:  "info",
		Output: os.Stdout,
	}
}
