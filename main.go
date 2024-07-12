package main

import (
	"os"
	"pp/api"
	"pp/database"
	"pp/router"
	"pp/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	var value string
	if len(os.Args) < 2 {
		value = "2"
	} else {
		// Read the argument
		value = os.Args[1]

	}
	length, _ := strconv.Atoi(value)

	redis_init := database.NewLRUCache_Redis()
	inmemory_init := database.NewLRUCache_Inmemory(1)
	usecase_init := usecase.Init(redis_init, inmemory_init)
	router_init := router.Init(usecase_init, length)
	api_init := api.Init(router_init)

	r := gin.Default()
	api_init.Application(r)
	r.Run()

}
