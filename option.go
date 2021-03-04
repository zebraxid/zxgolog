package zxgolog

import (
	"io"

	"github.com/zebraxid/zxgolog/logger"
)

// FnOption for optional functions when create new logger
type FnOption func(*option)

type option struct {
	loggerType LoggerType
	loggerOpt  *logger.Option
}

func defaultOption() *option {
	return &option{
		loggerType: Zerolog,
		loggerOpt:  logger.DefaultOption(),
	}
}

// LogType define which logger will be used, default is zerolog
func LogType(typ string) FnOption {
	return func(opt *option) {
		opt.loggerType = LoggerType(typ)
	}
}

// LogLevel for lowest level should be logged, default is info
func LogLevel(level string) FnOption {
	return func(opt *option) {
		if opt.loggerOpt == nil {
			return
		}

		opt.loggerOpt.Level = level
	}
}

// LogOutput for where log should be written, default is os.stdout
func LogOutput(output io.Writer) FnOption {
	return func(opt *option) {
		if opt.loggerOpt == nil {
			return
		}

		opt.loggerOpt.Output = output
	}
}

// AppName for set app name in log, if empty "app" won't be written in log
func AppName(name string) FnOption {
	return func(opt *option) {
		if opt.loggerOpt == nil {
			return
		}

		opt.loggerOpt.AppName = name
	}
}
