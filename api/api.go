package api

import (
	"pp/router"

	"github.com/gin-gonic/gin"
)

type Api struct {
	api *router.Router
}

func Init(a *router.Router) *Api {
	return &Api{
		api: a,
	}
}

func (a *Api) Application(r *gin.Engine) {

	r.GET("/print", a.api.Print)
	r.GET("/:key", a.api.GET)
	r.POST("/:key/:value/:time", a.api.Set)
	r.DELETE("/all", a.api.Del_all)
	r.DELETE("/:key", a.api.Del)

}
