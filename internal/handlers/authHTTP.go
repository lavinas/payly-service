package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lavinas/payly-service/internal/core/domains"
	"github.com/lavinas/payly-service/internal/core/ports"
)

type authHTTP struct {
	authService ports.AuthService
}

func NewauthHTTP(authService ports.AuthService) *authHTTP {
	return &authHTTP{authService: authService}
}

func mountError(code int, message string, description string) domains.AuthError {
	return domains.AuthError{
		Code:        code,
		Error:       message,
		Description: description,
	}
}

func throwError(c *gin.Context, err domains.AuthError) {
	c.IndentedJSON(err.Code, err)
}

func (a *authHTTP) Token(c *gin.Context) {
	var auth domains.AuthIn
	if err := c.BindJSON(&auth); err != nil {
		err := mountError(400, "invalid request", "invalid request")
		throwError(c, err)
	}
	t, err := a.authService.Token(auth)
	if err != nil {
		e := mountError(400, "invalid request", err.Error())
		throwError(c, e)
	}
	c.IndentedJSON(200, t)
}
