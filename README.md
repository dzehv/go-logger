# Golang logging package with levels

## Usage

``` golang
import log "github.com/dzehv/go-logger"

log.Info("Some info message")
log.Error("An error occured: %v", err)
log.Debug(0, "The sum is: %d", 1+1)
```
