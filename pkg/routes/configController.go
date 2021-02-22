package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wjdingman1/mediforui/pkg/config"
	e "github.com/wjdingman1/mediforui/pkg/error"
)

var ui *UI

// UI configuration constants for the client
type UI struct {
	DefaultFuserID string `json:"defaultFuserId"`
	EnableGroups   bool   `json:"enableGroups"`
	EnableDelete   bool   `json:"enableDelete"`
	UnknownUsers   string `json:"unknownUsers"`
	GroupPrefix    string `json:"groupPrefix"`
	UserTagPrefix  string `json:"userTagPrefix"`
	TagPrefixFlag  string `json:"tagPrefixFlag"`
}

// LoadConfig loads all routes associated with the configController
func LoadConfig(e *gin.Engine) {
	e.GET("/config", configUIHandler)
}

// This endpoint only returns the UI component of the configuration
func configUIHandler(c *gin.Context) {
	ui, err := newUI()
	log.Print("-- Retrieving UI configuration file --")
	if err != nil {
		e.HandleErrorResponse(c, err, 400)
	}
	c.JSON(http.StatusOK, ui)
}

// Unmarshall the ui config from the viper config object and return to client
func newUI() (*UI, error) {
	// If ui config has aready been unmarshalledm return it
	if ui != nil {
		return ui, nil
	}
	conf, err := config.New()
	if err != nil {
		return ui, err
	}
	if err = conf.UnmarshalKey(("UI"), &ui); err != nil {
		return ui, err
	}
	return ui, nil

}
