package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wjdingman1/mediforui/pkg/config"
	"github.com/wjdingman1/mediforui/pkg/error"
)

// LoadConfig loads all routes associated with the configController
func LoadConfig(e *gin.Engine) {
	e.GET("/config", configHandler)
}

func configHandler(c *gin.Context) {
	conf, err := config.New()
	log.Print("-- Retrieving configuration file --")
	if err != nil {
		error.HandleErrorResponse(c, err, 400)
	}
	c.JSON(http.StatusOK, conf.UI)
}
