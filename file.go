package config

import (
	"os"
	"path"
)

// LoadConfig loads the configuration from the specified path.
func LoadConfig(path string) (cfg IConfig, err error) {
	cfg = NewConfig(path)
	err = cfg.Load()
	return
}

// LoadAllConfigs loads all configuration files in the specified directory.
func LoadAllConfigs(name string) (configs []IConfig, err error) {
	dir, err := GetConfigDir(name)
	if err != nil {
		return
	}
	files, err := os.ReadDir(dir)
	if err != nil {
		return
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filePath := path.Join(dir, file.Name())
		cfg, err := LoadConfig(filePath)
		if err != nil {
			return nil, err
		}
		configs = append(configs, cfg)
	}
	return
}
