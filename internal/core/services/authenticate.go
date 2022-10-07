package services

import (
	"github.com/lavinas/payly-service/internal/core/domains"
	"github.com/lavinas/payly-service/internal/core/ports"
)

type authenticate struct {
	user ports.User
}

func NewAuthenticate(user ports.User) *authenticate {
	return &authenticate{user: user}
}

func (a *authenticate) Token(c ports.Context) {
	var auth domains.Auth
	if err := c.BindJSON(&auth); err != nil {
		arr := domains.AuthError{
			Error:       "invalid_request",
			Description: "invalid request",
		}
		c.IndentedJSON(400, arr)
		return
	}
	t := domains.Token{
		Code:   "5zVwuNnRDbF5YTJYDMfuZCnZtorxXdjjn9hssYiz",
		Type:   "Bearer",
		Expire: 57600,
	}
	c.IndentedJSON(200, t)
}
