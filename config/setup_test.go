package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_InstallSettings(t *testing.T) {
	dir, err := ioutil.TempDir("", "install_settings")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)

	cfg := DefaultConfig()
	cfg.ConfigDir = dir

	err = InstallSettings(cfg)
	assert.NoError(t, err)

	settingsFile := filepath.Join(cfg.ConfigDir, cfg.SettingsFile)

	_, err = os.Stat(settingsFile)
	assert.NoError(t, err)

	contents, err := ioutil.ReadFile(settingsFile)
	assert.NoError(t, err)

	assert.True(t, len(contents) > 0)
}
