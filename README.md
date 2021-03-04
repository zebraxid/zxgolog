# ZX Go Logger

> Simple and standardized logger to JSON as output.

## Getting Started

### Simple Logging Example

For simple logging, import the package **github.com/zebraxid/zxgolog**.

This method will write the log using default value.

```go
package main

import (
    "fmt"
    zxlog "github.com/zebraxid/zxgolog"
)

func main() {
    err := fmt.Errorf("err blabla")
    details := map[string]interface{}{"code": 123}
    argNum := 1

    zxlog.Logger().DebugDetail(details, "debug with detail")
    zxlog.Logger().Info("info has args %d", argNum)
    zxlog.Logger().Warn("warn only")
    zxlog.Logger().Error(err, "ouch error")
    zxlog.Logger().FatalDetail(err, details, "fatal, bye")
}

// Output: 
// {"level":"info","time":"2021-03-04T10:05:07+07:00","msg":"info has args 1"}
// {"level":"warn","time":"2021-03-04T10:05:07+07:00","msg":"warn only"}
// {"level":"error","time":"2021-03-04T10:05:07+07:00","err":"err blabla","msg":"ouch error"}
// {"level":"fatal","time":"2021-03-04T10:05:07+07:00","err":"err blabla","details":{"code":123},"msg":"fatal, bye"}
```
> Note: By default log level is Info, thus why Debug is not logged

### Initiate With Option Parameters

This option parameter will override default value for logging.

```go
package main

import (
    "fmt"
    zxlog "github.com/zebraxid/zxgolog"
)

func init() {
    if err := zxlog.Initiate(
        zxlog.AppName("zxapp"),
        zxlog.LogLevel("debug"),
    ); err != nil {
        // do something, better just exit the app
    }
}

func main() {
    zxlog.Logger().Debug("this debug will be logged")
}

// Output: 
// {"level":"debug","time":"2021-03-04T10:05:07+07:00","app":"zxapp","msg":"this debug will be logged"}
```
> Note: if use `zxlog.Initiate()` without option, this will initiate logger using default value.

### Available Options

#### AppName

```go
zxlog.Initiate(zxlog.AppName("zxapp"))
// Output: {...,"app":"zxapp"....}
```
> This option is to set "app" field in log text, used for application name. Default value is empty string.

#### LogLevel

```go
zxlog.Initiate(zxlog.LogLevel("error"))
// Output: only log with error level above will be logged
```
> This option is to set lowest level should be logged. Default value is "info".

#### LogOutput

```go
var someTextFile *os.File
zxlog.Initiate(zxlog.LogOutput(someTextFile))
// Output: log will be written in text file
```
> This option is to set target where log should be written. Default value is `os.Stdout`.

#### LogType

```go
zxlog.Initiate(zxlog.LogType("logrus"))
// Output: log will be written in logrus format
```
> This option is to set logger type will be used. Default value is zerolog.
