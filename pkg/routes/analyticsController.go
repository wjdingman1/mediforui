package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	e "github.com/wjdingman1/mediforui/pkg/error"
)

// AnalyticList describes a list of all the analytics running in the system
type AnalyticList struct {
	Analytics []Analytic `json:"analytics"`
}

// Analytic describes an analytic that is running in the system
type Analytic struct {
	Description string   `json:"description"`
	ID          string   `json:"id"`
	Media       []string `json:"media"`
	Name        string   `json:"name"`
}

// LoadAnalytics loads all the routers associated with the analyticController
func LoadAnalytics(e *gin.Engine) {
	e.GET("/analytics", analyticHandler)
}

// analyticHandler returns an HTTP response with the list of analytics running in the system
func analyticHandler(c *gin.Context) {
	analytics, err := newAnalyticList()
	log.Print("-- Retrieving Analytic List --")
	if err != nil {
		e.HandleErrorResponse(c, err, 400)
	} else {
		c.JSON(http.StatusOK, analytics)
	}
}

// newAnalyticList unmarshalls the analytic_config.json file into as struct and returns it
func newAnalyticList() (*AnalyticList, error) {
	a := new(AnalyticList)
	f, err := os.Open("analytic_config.json")
	if err != nil {
		return a, err
	}
	defer f.Close()
	bytes, _ := ioutil.ReadAll(f)
	if err := json.Unmarshal(bytes, &a); err != nil {
		return &AnalyticList{}, err
	}
	return a, nil

}
