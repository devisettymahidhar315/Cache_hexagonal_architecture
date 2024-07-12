package router

import (
	"net/http"
	"pp/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

var length = 2

type Router struct {
	d *usecase.Backends
}

func Init(w *usecase.Backends) *Router {
	return &Router{
		d: w,
	}
}

func (e *Router) Print(ctx *gin.Context) {
	output := e.d.Print()
	ctx.JSON(http.StatusOK, gin.H{
		"output": output,
	})
}

func (e *Router) Del_all(ctx *gin.Context) {
	e.d.Del_all()
	ctx.JSON(http.StatusOK, gin.H{
		"message": "deleted successfully",
	})
}

func (e *Router) Del(ctx *gin.Context) {
	key := ctx.Param("key")
	output := e.d.Del_key(key)
	ctx.JSON(http.StatusOK, gin.H{
		"output": output,
	})

}

func (e *Router) GET(ctx *gin.Context) {
	key := ctx.Param("key")
	output := e.d.Get(key)
	ctx.JSON(http.StatusOK, gin.H{
		"output": output,
	})

}

func (e *Router) Set(ctx *gin.Context) {
	key := ctx.Param("key")
	value := ctx.Param("value")
	time := ctx.Param("time")

	time1, _ := strconv.Atoi(time)

	e.d.Set(key, value, length, time1)
	ctx.JSON(http.StatusOK, gin.H{
		"output": "successfully key added",
	})
}
