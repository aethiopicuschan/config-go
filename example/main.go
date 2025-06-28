package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/aethiopicuschan/config-go"
)

func main() {
	// Set the user config directory to a known location
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Failed to get user home directory: %v", err)
	}
	configDirPath := filepath.Join(home, ".config")
	config.SetConfigDir(configDirPath)

	// Ensure the config directory exists
	dir, err := config.EnsureConfigDir("config-go-example")
	if err != nil {
		log.Fatalf("Failed to ensure config directory: %v", err)
	}

	// Create a new config file in the directory
	configPath := filepath.Join(dir, "config.json")
	conf := config.NewConfig(configPath)
	conf.Write([]byte("Hello World"))
	if err := conf.Save(); err != nil {
		log.Fatalf("Failed to save config: %v", err)
	}
}
