package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/aethiopicuschan/config-go"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	tmpDir := t.TempDir()

	tests := []struct {
		name string
		run  func(t *testing.T, cfg config.IConfig, filePath string)
	}{
		{
			name: "Write and Read buffer",
			run: func(t *testing.T, cfg config.IConfig, filePath string) {
				data := []byte("testdata")
				cfg.Write(data)

				buf, err := cfg.Read()
				assert.NoError(t, err)
				assert.Equal(t, data, buf)
				assert.Equal(t, cfg.Path(), filePath)

				// buffer should not be saved to file yet
				fileData, err := os.ReadFile(filePath)
				assert.True(t, os.IsNotExist(err))
				assert.Nil(t, fileData)
			},
		},
		{
			name: "Save and Load",
			run: func(t *testing.T, cfg config.IConfig, filePath string) {
				data := []byte("savecontent")
				cfg.Write(data)
				err := cfg.Save()
				assert.NoError(t, err)

				// clear buffer then load from file
				cfg.Write(nil)
				err = cfg.Load()
				assert.NoError(t, err)

				buf, err := cfg.Read()
				assert.NoError(t, err)
				assert.Equal(t, data, buf)
				assert.Equal(t, cfg.Path(), filePath)
			},
		},
		{
			name: "Reset content",
			run: func(t *testing.T, cfg config.IConfig, filePath string) {
				initial := []byte("initial")
				cfg.Write(initial)
				err := cfg.Save()
				assert.NoError(t, err)

				newContent := []byte("resetcontent")
				err = cfg.Reset(newContent)
				assert.NoError(t, err)

				// check buffer
				buf, err := cfg.Read()
				assert.NoError(t, err)
				assert.Equal(t, newContent, buf)

				// check file
				fileData, err := os.ReadFile(filePath)
				assert.NoError(t, err)
				assert.Equal(t, newContent, fileData)
				assert.Equal(t, cfg.Path(), filePath)
			},
		},
		{
			name: "Delete file",
			run: func(t *testing.T, cfg config.IConfig, filePath string) {
				content := []byte("somedata")
				cfg.Write(content)
				err := cfg.Save()
				assert.NoError(t, err)

				// confirm file exists
				_, err = os.Stat(filePath)
				assert.NoError(t, err)

				err = cfg.Delete()
				assert.NoError(t, err)

				_, err = os.Stat(filePath)
				assert.True(t, os.IsNotExist(err))
				assert.Equal(t, cfg.Path(), filePath)
			},
		},
		{
			name: "Read empty buffer",
			run: func(t *testing.T, cfg config.IConfig, filePath string) {
				_, err := cfg.Read()
				assert.ErrorIs(t, err, os.ErrInvalid)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			path := filepath.Join(tmpDir, tt.name+".json")
			cfg := config.NewConfig(path)

			tt.run(t, cfg, path)
		})
	}
}
