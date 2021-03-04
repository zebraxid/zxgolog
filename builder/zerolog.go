package builder

import (
	"github.com/zebraxid/zxgolog/logger"
	"github.com/zebraxid/zxgolog/logger/zerolog"
)

type zerologBuilder struct{}

// NewZerolog builder
func NewZerolog() LogBuilder {
	return new(zerologBuilder)
}

func (*zerologBuilder) Build(opt *logger.Option) (logger.Logger, error) {
	return zerolog.New(opt)
}
