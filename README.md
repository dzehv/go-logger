# Golang logging package with numeric levels

## Functions

- SetUTC() -- set UTC timezone flag
- SetDebugLvl(int) --  set debug level
- SetMsgPrefix(string) -- set all log messages prefix
- UseStderr(bool) -- if true log to stderr instead of stdout
- Debug(int, string, ...interface{}) -- log debug into including caller func name and line number
- Info(...)
- Warning(...)
- Error(...)
- Fatal(...)
- Panic(...) -- log panic message and than panic() with trace

## Usage

``` golang
import log "github.com/dzehv/go-logger"

log.Info("Some info message")
log.Error("An error occured: %v", err)
log.Debug(0, "The sum is: %d", 1+1)
```
