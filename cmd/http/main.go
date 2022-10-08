package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lavinas/payly-service/internal/core/services"
	"github.com/lavinas/payly-service/internal/handlers"
	"github.com/lavinas/payly-service/internal/repositories"
)

func main() {
	u := repositories.NewUserMysql()
	a := services.NewAuthenticate(u)
	h := handlers.NewauthHTTP(a)
	r := gin.Default()
	r.GET("/oauth/token", h.Token)
	r.Run("localhost:8080")
}
