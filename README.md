# nlog

A simple, pretty-printed JSON logging package for Go that provides structured logging with automatic trace information for errors.

## Features

- **Structured JSON logging** with pretty-printed output
- **Three log levels**: Info, Warn, and Error
- **Automatic trace capture** for errors (file, function, and line number)
- **Context support** for adding additional fields to log entries
- **Color-coded output** via go-prettyjson

## Installation

```bash
go get github.com/xDeFc0nx/nlog
```

## Usage

### Basic Logging

```go
package main

import "github.com/yourusername/nlog"

func main() {
    // Info logging
    nlog.Info("Application started", nil)

    // Info with context
    nlog.Info("User logged in", map[string]any{
        "user_id": 12345,
        "ip": "192.168.1.1",
    })

    // Warning
    nlog.Warn("API rate limit approaching", map[string]any{
        "remaining": 10,
        "limit": 100,
    })
}
```

### Error Logging with Automatic Trace

```go
func processData() error {
    err := someOperation()
    if err != nil {
        nlog.Error(err, map[string]any{
            "operation": "data_processing",
            "input_size": 1024,
        })
        return err
    }
    return nil
}
```

**Output:**

```json
{
  "level": "ERROR",
  "msg": "connection timeout",
  "trace": "main.go:42  main.processData",
  "operation": "data_processing",
  "input_size": 1024
}
```

## API Reference

### `Info(msg string, context map[string]any)`

Logs an informational message with optional context fields.

- **msg**: The log message
- **context**: Additional key-value pairs to include in the log entry (can be `nil`)

### `Warn(msg string, context map[string]any)`

Logs a warning message with optional context fields.

- **msg**: The log message
- **context**: Additional key-value pairs to include in the log entry (can be `nil`)

### `Error(err error, context map[string]any)`

Logs an error with automatic trace information (file, function, line number) and optional context fields.

- **err**: The error to log (if `nil`, nothing is logged)
- **context**: Additional key-value pairs to include in the log entry (can be `nil`)

## Dependencies

- [github.com/hokaccha/go-prettyjson](https://github.com/hokaccha/go-prettyjson) - For pretty-printed, color-coded JSON output

## License

MIT
