package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lavinas/payly-service/internal/core/ports"
	"github.com/lavinas/payly-service/internal/core/services"
	"github.com/lavinas/payly-service/internal/handlers"
	"github.com/lavinas/payly-service/internal/repositories"
	"github.com/lavinas/payly-service/internal/utils"
)

func main() {
	c := utils.NewConfig()
	u := repositories.NewUserMySQL(c)
	defer u.Close()
	t := utils.NewauthJWT(c)
	l := utils.NewlogFile(c, "auth")
	defer l.Close()
	a := services.NewAuthenticate(u, c, t, l)
	h := handlers.NewauthHTTP(a)
	r := ginConf(l)
	r.POST("/oauth/token", h.Token)
	srv := ginRun(l, r)
	ginShutDown(l, srv)
}

func ginConf(l ports.Log) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.MultiWriter(l.GetFile())
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	return r
}

func ginRun(l ports.Log, r http.Handler) *http.Server {
	l.Info("Starting GIN Service at 127.0.0.1:8081")
	srv := &http.Server{Addr: ":8081", Handler: r}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			l.Error("listen error: " + err.Error())
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	return srv
}

func ginShutDown(l ports.Log, srv *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		l.Error("Server Shutdown Error: " + err.Error())
	}
	<-ctx.Done()
	l.Info("GIN Server closed")
}
