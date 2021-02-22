package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/wjdingman1/mediforui/pkg/config"
	e "github.com/wjdingman1/mediforui/pkg/error"
)

//

// Facets configuration constants for the client
type Facets struct {
	Name        string `json:"name"`
	Description string `json:"Description"`
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

func newFacets() (*[]Facets, error) {
	var facets *[]Facets
	conf, err := config.New()
	if err != nil {
		return facets, err
	}

	// Viper.Get returns interface{}, assert to array of interface and assert element to map
	// https://stackoverflow.com/a/59759138
	for _, f := range conf.Get("FACETS").([]interface{}) {
		log.Print(f.(map[string]interface{})["name"])

	}

	return facets, nil

}
