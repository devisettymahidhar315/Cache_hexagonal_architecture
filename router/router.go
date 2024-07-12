package router

import (
	"net/http"
	"pp/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Router struct holds the backends and the cache length
type Router struct {
	d   *usecase.Backends
	len int
}

// Init initializes the Router with the given backends and cache length
func Init(w *usecase.Backends, length int) *Router {
	return &Router{
		d:   w,
		len: length,
	}
}

// Print handles the /print endpoint and returns the current state of the cache
func (e *Router) Print(ctx *gin.Context) {
	output := e.d.Print()
	ctx.JSON(http.StatusOK, gin.H{
		"output": output,
	})
}

// Del_all handles the /del_all endpoint and deletes all entries from the cache
func (e *Router) Del_all(ctx *gin.Context) {
	e.d.Del_all()
	ctx.JSON(http.StatusOK, gin.H{
		"message": "deleted successfully",
	})
}

// Del handles the /del/:key endpoint and deletes the specified key from the cache
func (e *Router) Del(ctx *gin.Context) {
	key := ctx.Param("key")
	output := e.d.Del_key(key)
	ctx.JSON(http.StatusOK, gin.H{
		"output": output,
	})
}

// GET handles the /get/:key endpoint and retrieves the value for the specified key from the cache
func (e *Router) GET(ctx *gin.Context) {
	key := ctx.Param("key")
	output := e.d.Get(key)
	ctx.JSON(http.StatusOK, gin.H{
		"output": output,
	})
}

// Set handles the /set/:key/:value/:time endpoint and sets the value for the specified key with an optional TTL
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
