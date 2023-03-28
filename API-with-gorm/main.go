package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const PORT = ":3000"

func main() {
	fmt.Println("Starting server.......")

	// db := database.Postgres()
	// controllers := controllers.New(db)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// // Users
	// users := r.Group("/v1/users")
	// {
	// 	users.POST("/", controllers.CreateUser)
	// 	users.GET("/products", controllers.GetUsersWithProducts)
	// }

	r.Run(PORT)
}
