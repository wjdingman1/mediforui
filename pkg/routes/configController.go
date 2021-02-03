package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wjdingman1/mediforui/pkg/config"
	e "github.com/wjdingman1/mediforui/pkg/error"
)

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
	log.Print("-- Retrieving configuration file --")
	if err != nil {
		e.HandleErrorResponse(c, err, 400)
	}
	c.JSON(http.StatusOK, ui)
}

// Create and return a new UI struct pointer
func newUI() (*UI, error) {
	conf, err := config.New()
	if err != nil {
		return &UI{}, err
	}
	return &UI{
		DefaultFuserID: conf.GetString("UI.DEFAULTFUSERID"),
		EnableGroups:   conf.GetBool("UI.ENABLEGROUPS"),
		EnableDelete:   conf.GetBool("UI.ENABLEDELETE"),
		UnknownUsers:   conf.GetString("UI.UNKNOWNUSERS"),
		GroupPrefix:    conf.GetString("UI.GROUPPREFIX"),
		UserTagPrefix:  conf.GetString("UI.USERTAGPREFIX"),
		TagPrefixFlag:  conf.GetString("UI.TAGPREFIXFLAG"),
	}, nil

}
