package config

import (
	"bytes"
	_ "embed"
	"github.com/spf13/viper"
	"os"
)

const (
	fileType = "yaml"
	appEnv = "APP_ENV"
)

//go:embed defaults/defaults.yaml
var defaultYamlFile []byte
//go:embed defaults/dev.yaml
var devYamlFile []byte

func LoadConfig(cfg *viper.Viper){
	currEnvironment, ok := os.LookupEnv(AppEnv.String())
	if !ok {
		currEnvironment = "dev"
	}
	cfg.SetConfigType(fileType)
	_ = cfg.ReadConfig(bytes.NewReader(defaultYamlFile))
	if currEnvironment == "dev" {
		_ = cfg.ReadConfig(bytes.NewReader(devYamlFile))
	}
}
