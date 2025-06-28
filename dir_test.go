package config_test

import (
	"os"
	"strings"
	"testing"

	"github.com/aethiopicuschan/config-go"
	"github.com/stretchr/testify/assert"
)

func TestConfigDirHelpers(t *testing.T) {
	tmpDir := t.TempDir()

	config.SetConfigDir(tmpDir)

	tests := []struct {
		name  string
		setup func(t *testing.T, dirPath, appName string)
		check func(t *testing.T, dirPath, appName string)
	}{
		{
			name: "DirExists returns false for non-existent directory",
			setup: func(t *testing.T, dirPath, appName string) {
				// do nothing
			},
			check: func(t *testing.T, dirPath, appName string) {
				exist, err := config.DirExists(appName)
				assert.NoError(t, err)
				assert.False(t, exist)
			},
		},
		{
			name: "DirExists returns true for existing directory",
			setup: func(t *testing.T, dirPath, appName string) {
				err := os.MkdirAll(dirPath, 0755)
				assert.NoError(t, err)
			},
			check: func(t *testing.T, dirPath, appName string) {
				exist, err := config.DirExists(appName)
				assert.NoError(t, err)
				assert.True(t, exist)
			},
		},
		{
			name: "EnsureConfigDir creates directory if missing",
			setup: func(t *testing.T, dirPath, appName string) {
				// do nothing
			},
			check: func(t *testing.T, dirPath, appName string) {
				dir, err := config.EnsureConfigDir(appName)
				assert.NoError(t, err)
				info, err := os.Stat(dir)
				assert.NoError(t, err)
				assert.True(t, info.IsDir())
			},
		},
		{
			name: "EnsureConfigDir does not recreate existing directory",
			setup: func(t *testing.T, dirPath, appName string) {
				err := os.MkdirAll(dirPath, 0755)
				assert.NoError(t, err)
			},
			check: func(t *testing.T, dirPath, appName string) {
				dir, err := config.EnsureConfigDir(appName)
				assert.NoError(t, err)
				info, err := os.Stat(dir)
				assert.NoError(t, err)
				assert.True(t, info.IsDir())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			appName := "myapp_" + strings.ReplaceAll(tt.name, " ", "_")

			dirPath, err := config.GetConfigDir(appName)

			assert.NoError(t, err)

			t.Cleanup(func() {
				os.RemoveAll(dirPath)
			})

			tt.setup(t, dirPath, appName)
			tt.check(t, dirPath, appName)
		})
	}
}
