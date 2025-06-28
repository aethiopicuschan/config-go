# config-go

[![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen?style=flat-square)](/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/aethiopicuschan/config-go.svg)](https://pkg.go.dev/github.com/aethiopicuschan/config-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/aethiopicuschan/config-go)](https://goreportcard.com/report/github.com/aethiopicuschan/config-go)
[![CI](https://github.com/aethiopicuschan/config-go/actions/workflows/ci.yaml/badge.svg)](https://github.com/aethiopicuschan/config-go/actions/workflows/ci.yaml)
[![codecov](https://codecov.io/gh/aethiopicuschan/config-go/graph/badge.svg)](https://codecov.io/gh/aethiopicuschan/config-go)

`config-go` is a Go library for handling configuration management in Go applications.

The configuration file can be automatically placed in an appropriate path depending on the operating system, or you can manually specify the path using the `SetConfigDir` function.

**Default paths:**

- **Linux:** `~/.config/appname/`
- **macOS:** `~/Library/Application Support/appname/`
- **Windows:** `%APPDATA%/appname/`

**SetConfigDir**

```go
home, _ := os.UserHomeDir()
configDirPath := filepath.Join(home, ".config")
config.SetConfigDir(configDirPath)
```

## Installation

```sh
go get -u github.com/aethiopicuschan/config-go
```

## Usage

The configuration file is read via the `IConfig` interface. Data is held as a byte slice in memory, and can be written to the file when needed.

```go
// IConfig defines the configuration interface
type IConfig interface {
	Save() error             // Save the buffered configuration to file
	Load() error             // Load the file content into buffer
	Reset(body []byte) error // Reset the configuration (buffer + file) to given default
	Delete() error           // Delete the configuration file
	Read() ([]byte, error)   // Return the current buffer content
	Write(body []byte)       // Overwrite buffer only (does not write to file immediately)
	Path() string            // Return the file path of the configuration
}
```

More details can be found in the [documentation](https://pkg.go.dev/github.com/aethiopicuschan/config-go).
You can also find reference implementations in the [example](example) directory.
