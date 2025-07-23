package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "Budi"},
	{ID: 2, Name: "Sari"},
}

func main() {
	_ = godotenv.Load()

	r := gin.Default()

	// GET /api/users
	r.GET("/api/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    users,
		})
	})

	// POST /api/users
	r.POST("/api/users", func(c *gin.Context) {
		var newUser struct {
			Name string `json:"name"`
		}

		if err := c.ShouldBindJSON(&newUser); err != nil || newUser.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Name is required",
			})
			return
		}

		user := User{
			ID:   len(users) + 1,
			Name: newUser.Name,
		}
		users = append(users, user)

		c.JSON(http.StatusCreated, gin.H{
			"success": true,
			"message": "User added",
			"data":    user,
		})
	})

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000" // Default port
	}
	fmt.Println("Server is running on port " + port)
	r.Run(":" + port)
}
