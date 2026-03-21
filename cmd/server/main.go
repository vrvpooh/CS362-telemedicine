package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"workflow-example.com/handler"
)

func main() {
	fmt.Println("Starting server")

	authHandler := handler.NewAuthHandler()

	r := gin.Default()

	r.POST("/api/auth/register", authHandler.Register)
	r.POST("/api/auth/login", authHandler.Login)
	r.GET("/api/users/me", authHandler.GetProfile)

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
