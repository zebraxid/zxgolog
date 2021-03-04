package builder

import (
	"zebrax.id/libs/zxgolog/logger"
	"zebrax.id/libs/zxgolog/logger/logrus"
)

type logrusBuilder struct{}

// NewLogrus builder
func NewLogrus() LogBuilder {
	return new(logrusBuilder)
}

func (*logrusBuilder) Build(opt *logger.Option) (logger.Logger, error) {
	return logrus.New(opt)
}
