package config

import (
	"os"
)

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

// Config holds the path and a buffer for configuration
type Config struct {
	path   string
	buffer []byte
}

// NewConfig creates a new Config with the given path
func NewConfig(path string) *Config {
	return &Config{
		path:   path,
		buffer: nil,
	}
}

// Save writes the current buffer to the file
func (c *Config) Save() error {
	return os.WriteFile(c.path, c.buffer, 0644)
}

// Load loads file content into the buffer
func (c *Config) Load() (err error) {
	data, err := os.ReadFile(c.path)
	if err != nil {
		return
	}
	c.buffer = data
	return
}

// Reset overwrites the buffer and the file with the given default
func (c *Config) Reset(body []byte) error {
	c.buffer = body
	return os.WriteFile(c.path, body, 0644)
}

// Delete removes the configuration file
func (c *Config) Delete() error {
	return os.Remove(c.path)
}

// Read returns the buffer content
func (c *Config) Read() (b []byte, err error) {
	if c.buffer == nil {
		err = os.ErrInvalid
		return
	}
	b = c.buffer
	return
}

// Write updates the buffer only (does not persist)
func (c *Config) Write(body []byte) {
	c.buffer = body
}

// Path returns the file path of the configuration
func (c *Config) Path() string {
	return c.path
}
