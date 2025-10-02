package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func createUser(c *gin.Context) {

}
func main() {
	router := gin.Default()
	router.POST("/api/users", createUser)
	fmt.Println("Hello World!!")
}
