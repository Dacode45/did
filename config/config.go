package config

import (
	"log"
	"os/user"
	"path/filepath"
)

// Config contains configuration paramaters for all of did
type Config struct {
	ConfigDir    string
	SettingsFile string
	TodoFile     string
	DoneFile     string
	DBFile       string
}

var rootConfig *Config

// SetConfig overrides the current config pramaters from the none
// empty fields of the config object
func SetConfig(c Config) {
	if c.ConfigDir != "" {
		rootConfig.ConfigDir = c.ConfigDir
	}

	if c.SettingsFile != "" {
		rootConfig.SettingsFile = c.SettingsFile
	}

	if c.TodoFile != "" {
		rootConfig.TodoFile = c.TodoFile
	}

	if c.DoneFile != "" {
		rootConfig.DoneFile = c.DoneFile
	}

	if c.DBFile != "" {
		rootConfig.DBFile = c.DBFile
	}
}

// GetConfig returns a clone of the current configuration
func GetConfig(c Config) Config {
	return *rootConfig
}

// DefaultConfig returns the configuration used at initialization
// but before the settings file is loaded
func DefaultConfig() Config {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	configDir := filepath.Join(user.HomeDir, ".did")

	return Config{
		ConfigDir:    configDir,
		SettingsFile: "settings.yaml",
		TodoFile:     "todo.txt",
		DoneFile:     "done.txt",
		DBFile:       "did.sqlite",
	}
}

// LoadDefaults sets the configation to the default
func LoadDefaults() {
	SetConfig(DefaultConfig())
}
