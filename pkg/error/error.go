package error

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// HandleErrorResponse calls PrintError and sends error status to client
func HandleErrorResponse(c *gin.Context, err error, status int) {
	PrintError(err)
	c.JSON(status, gin.H{"ERROR": errors.Wrap(err, "")})
}

// PrintError logs the error to the console
func PrintError(err error) {
	log.Printf("-- ERROR: %s --", err)
}
