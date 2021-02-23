package routes

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wjdingman1/mediforui/pkg/config"
	e "github.com/wjdingman1/mediforui/pkg/error"
)

// User struct defines the permissions, groups and names for the current user
type User struct {
	Admin       bool     `json:"admin"`
	DisplayName string   `json:"displayName"`
	Name        string   `json:"name"`
	Groups      []string `json:"groups"`
}

// LoadUser loads all the routes associated with the userController
func LoadUser(e *gin.Engine) {
	e.GET("/user", userHandler)
}

// userHandler returns an HTTP response with the User configuration info
func userHandler(c *gin.Context) {
	user, err := newUser(c)
	log.Print("-- Retrieving User configuration info --")
	if err != nil {
		e.HandleErrorResponse(c, err, 400)
	}
	c.JSON(http.StatusOK, user)

}

// extract the user information from the request headers and return
func newUser(c *gin.Context) (*User, error) {
	return &User{
		Admin:       checkAdmin(strings.Split(c.Request.Header.Get("groups"), ",")),
		DisplayName: string(c.Request.Header.Get("displayName")),
		Name:        string(c.Request.Header.Get("name")),
		Groups:      strings.Split(c.Request.Header.Get("groups"), ","),
	}, nil

}

// checkAdmin checks the request headers to see if the user is in the admin group
func checkAdmin(groups []string) bool {
	var formattedGroups []string
	if len(groups) == 0 {
		return false
	}
	config, err := config.New()
	if err != nil {
		return false
	}
	groupPrefix := config.Get("UI.GROUPPREFIX").(string)

	// format the groups to exclude the groupPrefix
	for _, group := range groups {
		if strings.HasPrefix(group, groupPrefix) {
			trimmedGroup := strings.Replace(group, groupPrefix, "", -1)
			formattedGroups = append(formattedGroups, trimmedGroup)
		}
	}
	// check if one of the formatted groups is admin
	for _, group := range formattedGroups {
		if group == "admin" || group == "Admin" {
			return true
		}
	}
	// if none of the groups are 'admin' return false
	return false

}
