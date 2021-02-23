package config

import (
	"strings"

	"github.com/spf13/viper"
)

func initViper() (*viper.Viper, error) {
	vip := viper.New()
	vip.AutomaticEnv()
	vip.SetConfigName("config")
	vip.SetConfigType("json")
	vip.AddConfigPath(".")
	// This allows for nested config values to be passed in as env vars with underscores
	// UI_ENABLEGROUPS -> UI.ENABLEGROUPS
	vip.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := vip.ReadInConfig(); err != nil {
		return vip, err
	}
	return vip, nil
}

// New returns a new viper configuration object
func New() (*viper.Viper, error) {
	vip, err := initViper()
	return vip, err
}
