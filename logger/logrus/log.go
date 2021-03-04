package logrus

import (
	"time"

	lr "github.com/sirupsen/logrus"
	"zebrax.id/libs/zxgolog/logger"
)

type logrusObj struct {
	*lr.Logger
	option *logger.Option
}

// New logrus logger
func New(opt *logger.Option) (logger.Logger, error) {
	lr.ErrorKey = "err"

	level, err := lr.ParseLevel(opt.Level)
	if err != nil {
		return nil, logger.ErrParseLevel(err)
	}

	lrLog := lr.New()
	lrLog.SetOutput(opt.Output)
	lrLog.SetLevel(level)
	lrLog.SetFormatter(new(lr.JSONFormatter))

	return &logrusObj{
		Logger: lrLog,
		option: opt,
	}, nil
}

func (l *logrusObj) Fatal(err error, msgFormat string, args ...interface{}) {
	entry := l.buildEntry(err, nil)
	if len(args) == 0 {
		entry.Fatal(msgFormat)
	} else {
		entry.Fatalf(msgFormat, args...)
	}
}

func (l *logrusObj) FatalDetail(err error, details map[string]interface{}, msgFormat string, args ...interface{}) {
	entry := l.buildEntry(err, details)
	if len(args) == 0 {
		entry.Log(lr.FatalLevel, msgFormat)
	} else {
		entry.Logf(lr.FatalLevel, msgFormat, args...)
	}
}

func (l *logrusObj) Error(err error, msgFormat string, args ...interface{}) {
	entry := l.buildEntry(err, nil)
	if len(args) == 0 {
		entry.Error(msgFormat)
	} else {
		entry.Errorf(msgFormat, args...)
	}
}

func (l *logrusObj) ErrorDetail(err error, details map[string]interface{}, msgFormat string, args ...interface{}) {
	entry := l.buildEntry(err, details)
	if len(args) == 0 {
		entry.Error(msgFormat)
	} else {
		entry.Errorf(msgFormat, args...)
	}
}

func (l *logrusObj) Warn(msgFormat string, args ...interface{}) {
	entry := l.buildEntry(nil, nil)
	if len(args) == 0 {
		entry.Warn(msgFormat)
	} else {
		entry.Warnf(msgFormat, args...)
	}
}

func (l *logrusObj) WarnDetail(details map[string]interface{}, msgFormat string, args ...interface{}) {
	entry := l.buildEntry(nil, details)
	if len(args) == 0 {
		entry.Warn(msgFormat)
	} else {
		entry.Warnf(msgFormat, args...)
	}
}

func (l *logrusObj) Info(msgFormat string, args ...interface{}) {
	entry := l.buildEntry(nil, nil)
	if len(args) == 0 {
		entry.Info(msgFormat)
	} else {
		entry.Infof(msgFormat, args...)
	}
}

func (l *logrusObj) InfoDetail(details map[string]interface{}, msgFormat string, args ...interface{}) {
	entry := l.buildEntry(nil, details)
	if len(args) == 0 {
		entry.Info(msgFormat)
	} else {
		entry.Infof(msgFormat, args...)
	}
}

func (l *logrusObj) Debug(msgFormat string, args ...interface{}) {
	entry := l.buildEntry(nil, nil)
	if len(args) == 0 {
		entry.Debug(msgFormat)
	} else {
		entry.Debugf(msgFormat, args...)
	}
}

func (l *logrusObj) DebugDetail(details map[string]interface{}, msgFormat string, args ...interface{}) {
	entry := l.buildEntry(nil, details)
	if len(args) == 0 {
		entry.Debug(msgFormat)
	} else {
		entry.Debugf(msgFormat, args...)
	}
}

func (l *logrusObj) buildEntry(err error, details map[string]interface{}) *lr.Entry {
	entry := l.Logger.WithTime(time.Now())

	if app := l.option.AppName; app != "" {
		entry = entry.WithField(logger.KeyApp, l.option.AppName)
	}

	if err != nil {
		entry = entry.WithError(err)
	}

	if details != nil {
		entry = entry.WithField(logger.KeyDetails, details)
	}

	return entry
}
