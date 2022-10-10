package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lavinas/payly-service/internal/core/services"
	"github.com/lavinas/payly-service/internal/handlers"
	"github.com/lavinas/payly-service/internal/repositories"
	"github.com/lavinas/payly-service/internal/utils"
	"io"
)

func main() {
	c := utils.NewConfig()
	u := repositories.NewUserMySQL(c)
	t := utils.NewauthJWT(c)
	l := utils.NewlogFile(c, "auth")
	a := services.NewAuthenticate(u, c, t, l)
	h := handlers.NewauthHTTP(a)
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.MultiWriter(l.GetFile())
	r := gin.Default()
	r.SetTrustedProxies([]string{"localhost"})
	r.POST("/oauth/token", h.Token)
	r.Run("localhost:8081")
}
