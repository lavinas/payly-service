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

func (a *authenticate) Token(domains.AuthIn) (domains.AuthToken, error) {
	t := domains.AuthToken{
		Code:   "5zVwuNnRDbF5YTJYDMfuZCnZtorxXdjjn9hssYiz",
		Type:   "Bearer",
		Expire: 57600,
	}
	return t, nil
}