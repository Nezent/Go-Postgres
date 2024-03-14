package main

import (
	"log"

	"github.com/Nezent/Go-PSQL/config"
	"github.com/Nezent/Go-PSQL/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	
	router := gin.Default()
	config.Connect()
	routes.UserRoute(router)
	router.Run("localhost:8080")
}