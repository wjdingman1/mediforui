package config

import (
	"github.com/spf13/viper"
)

// Config contains configuration values for the REST API
type Config struct {
	ContainerRoot   string
	WorkingDir      string
	WorkflowHost    string
	WorkflowPort    string
	CacheTTLSeconds int32
	Port            string
	UI              UI
	Facets          []Facets
}

// UI configuration constants for UI fields
type UI struct {
	DefaultFuserID string `json:"defaultFuserId"`
	EnableGroups   bool   `json:"enableGroups"`
	EnableDelete   bool   `json:"enableDelete"`
	UnknownUsers   string `json:"unknownUsers"`
	GroupPrefix    string `json:"groupPrefix"`
	UserTagPrefix  string `json:"userTagPrefix"`
	TagPrefixFlag  string `json:"tagPrefixFlag"`
}

// Facets are configuration values for Multi Class Ta2 algorithm
type Facets struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func initViper() (Config, error) {
	var config Config

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("../../pkg/config")
	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	return config, err
}

// New returns a new configuration object
func New() (*Config, error) {
	config, err := initViper()
	return &config, err
}
