package main

import (
	"pp/api"
	"pp/database"
	"pp/router"
	"pp/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	redis_init := database.NewLRUCache_Redis()
	inmemory_init := database.NewLRUCache_Inmemory(1)
	usecase_init := usecase.Init(redis_init, inmemory_init)
	router_init := router.Init(usecase_init)
	api_init := api.Init(router_init)

	r := gin.Default()
	api_init.Application(r)
	r.Run()

}
