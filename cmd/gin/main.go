package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lavinas/payly-service/internal/core/services"
	"github.com/lavinas/payly-service/internal/handlers"
	"github.com/lavinas/payly-service/internal/repositories"
	"github.com/lavinas/payly-service/internal/utils"
)

func main() {
	c := utils.NewConfig()
	u := repositories.NewUserMySQL(c)
	t := utils.NewauthJWT(c)
	a := services.NewAuthenticate(u, c, t)
	h := handlers.NewauthHTTP(a)
	r := gin.Default()
	r.SetTrustedProxies([]string{"localhost"})
	r.POST("/oauth/token", h.Token)
	r.Run("localhost:8081")
}
