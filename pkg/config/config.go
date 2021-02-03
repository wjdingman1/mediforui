package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	config Config
)

// Config contains configuration values for the REST API
type Config struct {
	ContainerRoot   string `mapstructure:"CONTAINER_ROOT"`
	WorkingDir      string `mapstructure:"WORKING_DIR"`
	WorkflowHost    string `mapstructure:"WORKFLOW_HOST"`
	WorkflowPort    string `mapstructure:"WORKFLOW_PORT"`
	CacheTTLSeconds int32  `mapstructure:"CACHE_TTL_SECONDS"`
	Port            string `mapstructure:"PORT"`
	UI              struct {
		DefaultFuserID string `json:"defaultFuserId"`
		EnableGroups   bool   `json:"enableGroups"`
		EnableDelete   bool   `json:"enableDelete"`
		UnknownUsers   string `json:"unknownUsers"`
		GroupPrefix    string `json:"groupPrefix"`
		UserTagPrefix  string `json:"userTagPrefix"`
		TagPrefixFlag  string `json:"tagPrefixFlag"`
	} `json:"UI"`
	Facets []struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"Facets"`
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
	viper.AutomaticEnv()
	viper.BindEnv("CONTAINER_ROOT")
	viper.BindEnv("WORKING_DIR")
	viper.BindEnv("WORKFLOW_HOST")
	viper.BindEnv("WORKFLOW_PORT")
	viper.BindEnv("CACHE_TTL_SECONDS")
	viper.BindEnv("groupPrefix")
	viper.BindEnv("PORT")

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}
	err := viper.Unmarshal(&config)

	log.Printf("%+v", config)
	return config, err
}

// New returns a new configuration object
func New() (*Config, error) {
	config, err := initViper()
	return &config, err
}
