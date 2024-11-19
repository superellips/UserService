package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/api/version/user/:userId", GetUser)
	router.POST("/api/version/users", PostUser)
	router.PUT("/api/version/users", PutUser)
	router.DELETE("/api/version/user/:userId", DeleteUser)

	router.Run(":8080")
}
