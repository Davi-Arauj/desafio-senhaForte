package main

import (
	"log"

	"github.com/desafio-senhaForte/api"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	api.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal()
	}
}
