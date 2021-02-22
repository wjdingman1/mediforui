package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/wjdingman1/mediforui/pkg/config"
	e "github.com/wjdingman1/mediforui/pkg/error"
)

var facets *[]Facets

// Facets configuration constants for the client
type Facets struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// LoadFacets loads all the routes associated with the facetsController
func LoadFacets(e *gin.Engine) {
	e.GET("/facets", configFacetsHandler)
}

func configFacetsHandler(c *gin.Context) {
	facets, err := newFacets()
	log.Print("-- Retrieving facets list --")
	if err != nil {
		e.HandleErrorResponse(c, err, 400)
	}
	c.JSON(http.StatusOK, facets)

}

// Unmarshall the facets from the viper config object and return to client
func newFacets() (*[]Facets, error) {
	// If facets has already been unmarshalled, return it
	if facets != nil {
		return facets, nil
	}
	conf, err := config.New()
	if err != nil {
		return facets, err
	}
	if err = conf.UnmarshalKey(("FACETS"), &facets); err != nil {
		return facets, err
	}
	return facets, nil
}
