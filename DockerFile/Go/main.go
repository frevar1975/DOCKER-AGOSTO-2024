package main

import (
	"github.com/gin-gonic/gin"
)

type Message struct {
	Message string `json:"message"`
}

func main() {
	r := gin.Default()
	r.GET("/", handleHelloWorld)
	r.Run(":8080")
}

func handleHelloWorld(c *gin.Context) {
	response := Message{
		Message: "Hello world",
	}
	c.JSON(200, response)
}
