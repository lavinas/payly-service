package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lavinas/payly-service/internal/core/domains"
	"github.com/lavinas/payly-service/internal/core/ports"
	"strings"
)

var (
	errorMap = map[string]int{
		"invalid_request":     400,
		"internal error":      500,
		"invalid_client":      401,
		"invalid_credentials": 401,
	}
)

type authHTTP struct {
	authService ports.AuthService
}

func NewauthHTTP(authService ports.AuthService) *authHTTP {
	return &authHTTP{authService: authService}
}

func (a *authHTTP) Token(c *gin.Context) {
	var auth domains.AuthIn
	if err := c.Bind(&auth); err != nil {
		throwError(c, "invalid_request: "+err.Error())
		return
	}
	t, err := a.authService.Token(auth)
	if err != nil {
		throwError(c, err.Error())
		return
	}
	c.JSON(200, t)
}

func throwError(c *gin.Context, message string) {
	sp := strings.Split(message, ": ")
	if (len(sp)) != 2 {
		c.JSON(500, domains.AuthError{
			Error:       "internal error",
			Description: "plese contact Vooo admin"},
		)
		return
	}
	code, ok := errorMap[sp[0]]
	if !ok {
		c.JSON(500, domains.AuthError{
			Error:       "internal error",
			Description: "plese contact Vooo admin"},
		)
		return
	}
	c.JSON(code, domains.AuthError{
		Code:        code,
		Error:       sp[0],
		Description: sp[1],
	})
}
