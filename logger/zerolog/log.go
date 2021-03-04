package zerolog

import (
	zl "github.com/rs/zerolog"
	"github.com/zebraxid/zxgolog/logger"
)

type zerologObj struct {
	*zl.Logger
	option *logger.Option
}

// New zerolog logger
func New(opt *logger.Option) (logger.Logger, error) {
	zl.ErrorFieldName = "err"
	zl.MessageFieldName = "msg"

	// override this because zerolog parser doesn't recognize warning
	if opt.Level == "warning" {
		opt.Level = "warn"
	}

	level, err := zl.ParseLevel(opt.Level)
	if err != nil {
		return nil, logger.ErrParseLevel(err)
	}

	zlLog := zl.New(opt.Output).Level(level)
	return &zerologObj{
		Logger: &zlLog,
		option: opt,
	}, nil
}

func (l *zerologObj) Fatal(err error, msgFormat string, args ...interface{}) {
	event := l.buildEvent(l.Logger.Fatal(), err, nil)
	write(event, msgFormat, args...)
}

func (l *zerologObj) FatalDetail(err error, details map[string]interface{}, msgFormat string, args ...interface{}) {
	event := l.buildEvent(l.Logger.Fatal(), err, details)
	write(event, msgFormat, args...)
}

func (l *zerologObj) Error(err error, msgFormat string, args ...interface{}) {
	event := l.buildEvent(l.Logger.Error(), err, nil)
	write(event, msgFormat, args...)
}

func (l *zerologObj) ErrorDetail(err error, details map[string]interface{}, msgFormat string, args ...interface{}) {
	event := l.buildEvent(l.Logger.Error(), err, details)
	write(event, msgFormat, args...)
}

func (l *zerologObj) Warn(msgFormat string, args ...interface{}) {
	event := l.buildEvent(l.Logger.Warn(), nil, nil)
	write(event, msgFormat, args...)
}

func (l *zerologObj) WarnDetail(details map[string]interface{}, msgFormat string, args ...interface{}) {
	event := l.buildEvent(l.Logger.Warn(), nil, details)
	write(event, msgFormat, args...)
}

func (l *zerologObj) Info(msgFormat string, args ...interface{}) {
	event := l.buildEvent(l.Logger.Info(), nil, nil)
	write(event, msgFormat, args...)
}

func (l *zerologObj) InfoDetail(details map[string]interface{}, msgFormat string, args ...interface{}) {
	event := l.buildEvent(l.Logger.Info(), nil, details)
	write(event, msgFormat, args...)
}

func (l *zerologObj) Debug(msgFormat string, args ...interface{}) {
	event := l.buildEvent(l.Logger.Debug(), nil, nil)
	write(event, msgFormat, args...)
}

func (l *zerologObj) DebugDetail(details map[string]interface{}, msgFormat string, args ...interface{}) {
	event := l.buildEvent(l.Logger.Debug(), nil, details)
	write(event, msgFormat, args...)
}

func (l *zerologObj) buildEvent(levelEvent *zl.Event, err error, details map[string]interface{}) *zl.Event {
	event := levelEvent.Timestamp()

	if app := l.option.AppName; app != "" {
		event = event.Str(logger.KeyApp, app)
	}

	if err != nil {
		event = event.Err(err)
	}

	if details != nil {
		event = event.Interface(logger.KeyDetails, details)
	}

	return event
}

func write(event *zl.Event, msgFormat string, args ...interface{}) {
	if len(args) == 0 {
		event.Msg(msgFormat)
	} else {
		event.Msgf(msgFormat, args...)
	}
}
