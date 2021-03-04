package builder

import (
	"zebrax.id/libs/zxgolog/logger"
)

// LogBuilder for build logger instance
type LogBuilder interface {
	Build(opt *logger.Option) (logger.Logger, error)
}
