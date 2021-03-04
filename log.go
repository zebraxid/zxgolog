package zxgolog

import (
	"sync"

	"github.com/pkg/errors"
	"zebrax.id/libs/zxgolog/builder"
	"zebrax.id/libs/zxgolog/logger"
)

// LoggerType ...
type LoggerType string

const (
	//Logrus type
	Logrus LoggerType = "logrus"
	//Zerolog type
	Zerolog LoggerType = "zerolog"
)

var (
	mapBuilder map[LoggerType]builder.LogBuilder
	gLogger    logger.Logger
	onceLogger sync.Once
)

func init() {
	mapBuilder = map[LoggerType]builder.LogBuilder{
		Logrus:  builder.NewLogrus(),
		Zerolog: builder.NewZerolog(),
	}
}

// New for initiate new logger based on logger type
func New(opts ...FnOption) (logger.Logger, error) {
	option := defaultOption()
	for _, opt := range opts {
		opt(option)
	}

	logBuilder, _ := mapBuilder[option.loggerType]
	if logBuilder == nil {
		return nil, errors.Errorf("logger type [%s] is not implemented", option.loggerType)
	}

	return logBuilder.Build(option.loggerOpt)
}

// Initiate for initiate new global logger based on logger type, can be called one time only
func Initiate(opts ...FnOption) error {
	if gLogger != nil {
		return errors.New("logger already initiated")
	}

	var err error
	gLogger, err = New(opts...)
	return err
}

// Logger is global logger from first initiate
func Logger() logger.Logger {
	onceLogger.Do(func() {
		Initiate()
	})

	return gLogger
}
