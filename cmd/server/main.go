package main

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wjdingman1/mediforui/pkg/config"
	"github.com/wjdingman1/mediforui/pkg/routes"
)

func main() {
	conf, err := config.New()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	routes.LoadConfig(r)
	routes.LoadFacets(r)
	r.Run(strings.Join([]string{":", conf.Port}, ""))
}
