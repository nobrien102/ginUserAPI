package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []user

func main() {
	r := gin.Default()

	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", getUsers)
		userRoutes.POST("/", CreateUser)
		userRoutes.PUT("/:id", EditUser)
		userRoutes.DELETE("/:id", DeleteUser)

	}

	if err := r.Run(":5000"); err != nil {
		log.Fatal(err.Error())
	}
}

func getUsers(c *gin.Context) {
	c.JSON(200, users)
}

func CreateUser(c *gin.Context) {
	var reqBody user
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
		return
	}

	reqBody.ID = uuid.New().String()

	users = append(users, reqBody)

	c.JSON(200, gin.H{
		"error": false,
	})
}

func EditUser(c *gin.Context) {
	id := c.Param("id")
	var reqBody user
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
		return
	}

	for i, u := range users {
		if u.ID == id {
			users[i].Name = reqBody.Name
			users[i].Age = reqBody.Age

			c.JSON(200, gin.H{
				"error": false,
			})
			return
		}
	}

	c.JSON(404, gin.H{
		"error":   true,
		"message": "could not find user with supplied id",
	})

}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)

			c.JSON(200, gin.H{
				"error": false,
			})
			return
		}
	}

	c.JSON(404, gin.H{
		"error":   true,
		"message": "Id not found",
	})
}
