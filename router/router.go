package router

import (
	"net/http"
	"pp/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Router struct {
	d   *usecase.Backends
	len int
}

func Init(w *usecase.Backends, length int) *Router {
	return &Router{
		d:   w,
		len: length,
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

	e.d.Set(key, value, e.len, time1)
	ctx.JSON(http.StatusOK, gin.H{
		"output": "successfully key added",
	})
}
