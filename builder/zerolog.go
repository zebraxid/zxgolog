package builder

import (
	"zebrax.id/libs/zxgolog/logger"
	"zebrax.id/libs/zxgolog/logger/zerolog"
)

type zerologBuilder struct{}

// NewZerolog builder
func NewZerolog() LogBuilder {
	return new(zerologBuilder)
}

func (*zerologBuilder) Build(opt *logger.Option) (logger.Logger, error) {
	return zerolog.New(opt)
}
