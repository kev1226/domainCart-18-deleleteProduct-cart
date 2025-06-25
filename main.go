package main

import (
	"deleteProduct-cart/config"
	"deleteProduct-cart/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectRedis()
	r := gin.Default()
	routes.SetupRoutes(r)

	if err := r.Run(":3033"); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
