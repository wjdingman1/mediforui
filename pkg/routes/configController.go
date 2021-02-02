package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wjdingman1/mediforui/pkg/config"
)

// LoadConfig loads all routes associated with the configController
func LoadConfig(e *gin.Engine) {
	e.GET("/config", configHandler)
}

func configHandler(c *gin.Context) {
	conf, err := config.New()
	log.Print("-- Retrieving configuration file --")
	if err != nil {
		log.Printf("ERROR retrieving configuration file %s", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, conf.UI)
}
