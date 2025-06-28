package config_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/aethiopicuschan/config-go"
	"github.com/stretchr/testify/assert"
)

func TestLoadAllConfigs(t *testing.T) {
	tmpDir := t.TempDir()

	config.SetConfigDir(tmpDir)

	tests := []struct {
		name           string
		setupFiles     map[string]string
		expectErr      bool
		expectedConfig map[string]string
	}{
		{
			name: "loads multiple configs",
			setupFiles: map[string]string{
				"config1.json": `{"a":1}`,
				"config2.json": `{"b":2}`,
			},
			expectErr: false,
			expectedConfig: map[string]string{
				"config1.json": `{"a":1}`,
				"config2.json": `{"b":2}`,
			},
		},
		{
			name: "returns error on unreadable file",
			setupFiles: map[string]string{
				"broken.json": "",
			},
			expectErr: true,
		},
		{
			name:           "directory is empty",
			setupFiles:     nil,
			expectErr:      false,
			expectedConfig: map[string]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Generate a unique appName from the test name
			appName := "myapp_" + strings.ReplaceAll(tt.name, " ", "_")

			// Create a test directory
			confDir, err := config.EnsureConfigDir(appName)
			assert.NoError(t, err)

			// Clean up after the test
			t.Cleanup(func() {
				os.RemoveAll(confDir)
			})

			// Create test files
			if tt.setupFiles != nil {
				for fname, body := range tt.setupFiles {
					fullPath := filepath.Join(confDir, fname)
					err := os.WriteFile(fullPath, []byte(body), 0644)
					assert.NoError(t, err)
				}
			}

			// Make the file unreadable
			if tt.name == "returns error on unreadable file" {
				brokenFile := filepath.Join(confDir, "broken.json")
				err := os.Chmod(brokenFile, 0000)
				assert.NoError(t, err)
				t.Cleanup(func() {
					os.Chmod(brokenFile, 0644)
				})
			}

			configs, err := config.LoadAllConfigs(appName)

			if tt.expectErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, len(tt.expectedConfig), len(configs))

			for _, cfg := range configs {
				read, readErr := cfg.Read()
				assert.NoError(t, readErr)
				filename := filepath.Base(cfg.(*config.Config).Path())

				expected, ok := tt.expectedConfig[filename]
				assert.True(t, ok, "unexpected file: %s", filename)
				assert.Equal(t, expected, string(read))
			}
		})
	}
}
