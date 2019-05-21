package myhttp

import "gopkg.in/ini.v1"

type ConfigModel struct {
	Header map[string]string
}

func LoadByCfg(file *ini.File) *ConfigModel {
	model := new(ConfigModel)
	model.Header = map[string]string{}
	headerSection := file.Section("Header")
	for key, value := range headerSection.KeysHash() {
		model.Header[key] = value
	}
	return model
}
