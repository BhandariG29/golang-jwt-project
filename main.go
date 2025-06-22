package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"

	routes "github.com/BhandariG29/golang-jwt-project/routes"
)

func main(){
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("error in loading dotenv file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context){
		c.JSON(200, gin.H{"success":"access granted for api-1"})
	})
	router.GET("/api-2", func(c *gin.Context){
		c.JSON(200, gin.H{"success":"access granted for api-2"})
	})

	router.Run(":" + port)
}