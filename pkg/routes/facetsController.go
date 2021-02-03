package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/wjdingman1/mediforui/pkg/config"
	"github.com/wjdingman1/mediforui/pkg/error"
)

// LoadFacets loads all the routes associated with the facetsController
func LoadFacets(e *gin.Engine) {
	e.GET("/facets", facetsHandler)
}

func facetsHandler(c *gin.Context) {
	conf, err := config.New()
	log.Print("-- Retrieving facets list --")
	if err != nil {
		error.HandleErrorResponse(c, err, 400)
	}
	c.JSON(http.StatusOK, conf)

}
