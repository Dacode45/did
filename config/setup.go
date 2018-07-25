package config

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

func init() {
	LoadTemplates()
}

// InitializeDid Takes the current configuration, and creates
// the directories and files. Will also take care of problems
// Caused by crashing.
func InitializeDid(cfg Config) error {
	if rootConfig == nil {
		return fmt.Errorf("No configuation loaded. Was SetConfig called?")
	}

	// Check that folder exist
	dirStat, err := os.Stat(rootConfig.ConfigDir)
	if os.IsNotExist(err) {
		err = SetupFilesAndFolders(cfg)
		if err != nil {
			return err
		}
	}
	// make sure it is a directory
	if !dirStat.IsDir() {
		return fmt.Errorf("Initialization Error: config dir %q is not a directory", rootConfig.ConfigDir)
	}

	//

	return nil
}

// SetupFilesAndFolders creates configuration files for did
func SetupFilesAndFolders(cfg Config) error {
	err := os.MkdirAll(rootConfig.ConfigDir, os.ModePerm)
	if err != nil {
		return err
	}
	return err
}

// InstallSettings installs the settings file
func InstallSettings(cfg Config) error {
	settings := SettingsTempl{
		Version: Version,
		Conf:    cfg,
	}
	settingsFile, err := templateBox.MustString("settings.yaml.tmpl")
	if err != nil {
		return err
	}
	tmpl, err := template.New("settings").Parse(settingsFile)
	if err != nil {
		return err
	}

	f, err := os.Create(filepath.Join(cfg.ConfigDir, cfg.SettingsFile))
	if err != nil {
		return err
	}
	defer f.Close()

	err = tmpl.Execute(f, settings)
	if err != nil {
		return err
	}

	return nil
}
