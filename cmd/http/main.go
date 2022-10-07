package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lavinas/payly-service/internal/handlers"
	"github.com/lavinas/payly-service/internal/core/services"
)

func main() {
    user := handlers.NewUserMysql()
    auth := services.NewAuthenticate(user)
	router := gin.Default()
	router.GET("/oauth/token", auth.Token)
	router.Run("localhost:8080")
}
