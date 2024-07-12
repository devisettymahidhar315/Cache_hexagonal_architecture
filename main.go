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
		// Default value for cache length if no argument is provided
		value = "2"
	} else {
		// Read the argument from the command line
		value = os.Args[1]
	}
	// Convert the argument to an integer
	length, _ := strconv.Atoi(value)

	// Initialize Redis LRU Cache
	redis_init := database.NewLRUCache_Redis()
	// Initialize In-memory LRU Cache
	inmemory_init := database.NewLRUCache_Inmemory(1)
	// Initialize use case with both caches
	usecase_init := usecase.Init(redis_init, inmemory_init)
	// Initialize router with use case and cache length
	router_init := router.Init(usecase_init, length)
	// Initialize API with the router
	api_init := api.Init(router_init)

	// Create a new Gin engine
	r := gin.Default()
	// Setup routes and handlers
	api_init.Application(r)
	// Start the HTTP server
	r.Run()
}
