package builder

import (
	"github.com/zebraxid/zxgolog/logger"
)

// LogBuilder for build logger instance
type LogBuilder interface {
	Build(opt *logger.Option) (logger.Logger, error)
}
