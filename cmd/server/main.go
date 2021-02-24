package main

import (
	"log"
	"os"
	"strings"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/wjdingman1/mediforui/pkg/config"
	"github.com/wjdingman1/mediforui/pkg/directories"
	"github.com/wjdingman1/mediforui/pkg/routes"
)

func main() {
	conf, err := config.New()
	if err != nil {
		panic("-- CONFIG NOT FOUND: SHUTTING DOWN --")
	}
	r := gin.Default()
	routes.LoadUser(r)
	routes.LoadConfig(r)
	routes.LoadFacets(r)
	routes.LoadAnalytics(r)
	r.Run(strings.Join([]string{":", conf.GetString("PORT")}, ""))
}

// create application directories for storing uploads
func init() {
	applicationDirectories := directories.New().ApplicationDirectory
	appMap := structs.Map(applicationDirectories)
	for _, directory := range appMap {
		if err := os.MkdirAll(directory.(string), 0775); err != nil {
			log.Printf("Error creating directory - %s", directory.(string))
		} else {
			log.Printf("Creating application directory - %s", directory.(string))
		}
	}
}
