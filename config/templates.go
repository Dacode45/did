package config

import "github.com/gobuffalo/packr"

var templateBox packr.Box

// LoadTemplates loads the template files from mem
func LoadTemplates() {
	templateBox = packr.NewBox("./templates")
}

// SettingsTempl contains the fields needed to populate the settings.yaml
// file
type SettingsTempl struct {
	Version string
	Conf    Config
}
