package logger

import "github.com/pkg/errors"

const (
	// KeyDetails for field "details" in json
	KeyDetails = "details"
	// KeyApp for field "app" in json
	KeyApp = "app"
)

// Logger provide common logging functions
type Logger interface {
	Fatal(err error, msgFormat string, args ...interface{})
	FatalDetail(err error, details map[string]interface{}, msgFormat string, args ...interface{})
	Error(err error, msgFormat string, args ...interface{})
	ErrorDetail(err error, details map[string]interface{}, msgFormat string, args ...interface{})
	Warn(msgFormat string, args ...interface{})
	WarnDetail(details map[string]interface{}, msgFormat string, args ...interface{})
	Info(msgFormat string, args ...interface{})
	InfoDetail(details map[string]interface{}, msgFormat string, args ...interface{})
	Debug(msgFormat string, args ...interface{})
	DebugDetail(details map[string]interface{}, msgFormat string, args ...interface{})
}

// ErrParseLevel ...
func ErrParseLevel(err error) error {
	return errors.Wrap(err, "parsing level")
}
