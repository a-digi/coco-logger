# coco-logger

A lightweight file logger for Go with daily folder structure and simple API.

## Features

- Writes to log files with automatic, date-based folder structure (YYYY/MM/DD)
- Thread-safe logging (mutex-protected)
- Simple methods: Info, Warning, Error, Close
- Configurable log directory path

## Installation

Use Go Modules and add the module to your project. Adjust the module path to your actual repository.

```bash
go get github.com/your-org/coco-logger
```

Import the logger from the `src/logger` package:

```go
import logger "github.com/your-org/coco-logger/src/logger"
```

## Quick Start

```go
package main

import (
    logger "github.com/your-org/coco-logger/src/logger"
)

func main() {
    // Set your log directory (structured by date)
    log, err := logger.NewLogger("app.log", "./logs")
    if err != nil {
        panic(err)
    }
    defer log.Close()

    log.Info("service starting: version %s", "0.1.0")
    log.Warning("slower than expected: duration_ms=%d", 120)
    log.Error("failed to do work: err=%s", "some error")
}
```

Logs are written to a path like `./logs/2026/02/14/app.log`.

## API

- `NewLogger(fileName string, logDir string) (Logger, error)`
  - Creates a new logger that writes to `logDir/YYYY/MM/DD/fileName`
  - `logDir` must be set (no fallback)
- `Info(msg string, args ...interface{})`
- `Warning(msg string, args ...interface{})`
- `Error(msg string, args ...interface{})`
- `Close()` closes the file resource

### Format

Each log entry is written as a simple text line:

```
[YYYY-MM-DD HH:MM:SS] [LEVEL] message\n
```

Example:

```
[2026-02-14 10:23:45] [INFO] service starting: version 0.1.0
```

## Best Practices

- Always call `Close()` (e.g., with `defer`) to close file handles
- Choose a meaningful `logDir` (e.g., `./logs` or `/var/log/myapp`)
- Output structured information as part of the message (key=value), e.g., `duration_ms=120`
- Avoid logging sensitive data (PII/secrets)

## Tests

If tests are present, run them with:

```bash
go test ./...
```

## Versioning

This project follows semantic versioning; breaking changes may occur until v1.0.

## License

This project is released under the MIT license. See `LICENSE`.
