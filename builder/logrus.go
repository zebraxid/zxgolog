package builder

import (
	"github.com/zebraxid/zxgolog/logger"
	"github.com/zebraxid/zxgolog/logger/logrus"
)

type logrusBuilder struct{}

// NewLogrus builder
func NewLogrus() LogBuilder {
	return new(logrusBuilder)
}

func (*logrusBuilder) Build(opt *logger.Option) (logger.Logger, error) {
	return logrus.New(opt)
}
