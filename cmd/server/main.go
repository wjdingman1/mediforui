package main

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wjdingman1/mediforui/pkg/routes"
)

func main() {

	r := gin.Default()
	routes.LoadConfig(r)
	routes.LoadFacets(r)
	r.Run(strings.Join([]string{":", "3000"}, ""))
}
