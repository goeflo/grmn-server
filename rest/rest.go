package rest

import (
	"fmt"
	"grmn-server/activities"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// RestOpts all options to start rest api server
type RestOpts struct {
	Activities string
	Port       int
	Verbose    bool
}

var opts = RestOpts{}

// Start start the rest server
func Start(startOpts RestOpts) {
	opts = startOpts
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/activities", getActivities)
	router.GET("/activity/:name", getActivity)
	router.Run(fmt.Sprintf("localhost:%v", opts.Port))
}

func getActivities(c *gin.Context) {
	resp := activities.GetListOfActivities(opts.Activities)
	if opts.Verbose {
		fmt.Printf("return fit list size: %v\n", len(resp))
	}
	c.JSON(http.StatusOK, resp)
}

func getActivity(c *gin.Context) {

}
