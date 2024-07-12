package api

import (
	"pp/router"

	"github.com/gin-gonic/gin"
)

// Api struct holds the router
type Api struct {
	api *router.Router
}

// Init initializes the Api with the given router
func Init(a *router.Router) *Api {
	return &Api{
		api: a,
	}
}

// Application sets up the HTTP routes for the API
func (a *Api) Application(r *gin.Engine) {
	// Route to print the current state of the cache
	r.GET("/print", a.api.Print)
	// Route to get the value for a specified key
	r.GET("/:key", a.api.GET)
	// Route to set a value for a specified key with an optional TTL
	r.POST("/:key/:value/:time", a.api.Set)
	// Route to delete all entries from the cache
	r.DELETE("/all", a.api.Del_all)
	// Route to delete a specified key from the cache
	r.DELETE("/:key", a.api.Del)
}
