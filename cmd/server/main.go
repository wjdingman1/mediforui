package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wjdingman1/mediforui/pkg/routes"
)

func main() {
	r := gin.Default()
	routes.LoadConfig(r)
	r.Run(":3000")
}
