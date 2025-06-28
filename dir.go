package config

import (
	"os"
	"path"
)

var userConfigDir = os.UserConfigDir

// SetConfigDir allows overriding the function that returns the user configuration directory.
func SetConfigDir(dir string) {
	userConfigDir = func() (string, error) {
		return dir, nil
	}
}

// GetConfigDir returns the configuration directory for the given name.
func GetConfigDir(name string) (dir string, err error) {
	base, err := userConfigDir()
	if err != nil {
		return
	}
	dir = path.Join(base, name)
	return
}

// DirExists checks if the configuration directory for the given name exists.
// DirExists checks if the configuration directory for the given name exists.
func DirExists(name string) (exist bool, err error) {
	dir, err := GetConfigDir(name)
	if err != nil {
		return
	}
	info, statErr := os.Stat(dir)
	if os.IsNotExist(statErr) {
		exist = false
		err = nil
		return
	}
	if statErr != nil {
		exist = false
		err = statErr
		return
	}
	if info.IsDir() {
		exist = true
	} else {
		exist = false
		err = os.ErrInvalid
	}
	return
}

// EnsureConfigDir ensures that the configuration directory for the given name exists.
func EnsureConfigDir(name string) (dir string, err error) {
	dir, err = GetConfigDir(name)
	if err != nil {
		if err != os.ErrNotExist {
			return
		}
	}
	exist, err := DirExists(name)
	if err != nil {
		return
	}
	if !exist {
		err = os.MkdirAll(dir, 0755)
		return
	}
	return
}
