package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json: username`
	Password string `json: password`
}

func setupRouter() *gin.Engine {

	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "REST API")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		c.JSON(http.StatusOK, gin.H{"user": user})
	})

	// Get some JSON value
	r.GET("/getJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"user": "saad",
			"language":  "golang",
		}
		c.AsciiJSON(http.StatusOK, data)
	})

	//using POST method, post some JSON and get back
	r.POST("/postdata", func(c *gin.Context) {
		var u User
		c.BindJSON(&u)
		c.JSON(http.StatusOK, gin.H{
			"user": u.Username,
			"pass": u.Password,
		})
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}