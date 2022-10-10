package utils

import (
	"github.com/kataras/jwt"
	"github.com/lavinas/payly-service/internal/core/ports"
	"strconv"
	"time"
)

type claim struct {
	ClientId    string `json:"client_id"`
	Username    string `json:"username"`
	UserId      int    `json:"user_id"`
	GenetrateAt int64  `json:"generate_at"`
	ExpiresAt   int64  `json:"expires_at"`
}

type authJWT struct {
	lag int64
	key string
}

func NewauthJWT(config ports.Config) *authJWT {
	e := gerParam(config, "expires_in")
	lag, err := strconv.ParseInt(e, 10, 64)
	if err != nil {
		panic("auth configuration error: experes_in not numeric")
	}
	key := gerParam(config, "key")
	return &authJWT{lag: lag, key: key}
}

func (a *authJWT) Get(clientId string, username string, userId int) (string, int64, error) {
	g := time.Now().Unix()
	e := g + a.lag
	c := claim{ClientId: clientId, Username: username, UserId: userId, GenetrateAt: g, ExpiresAt: e}
	t, err := jwt.Sign(jwt.HS256, []byte(a.key), c)
	if err != nil {
		return "", 0, err
	}
	return string(t), a.lag, nil
}

func gerParam(config ports.Config, index string) string {
	e, err := config.GetField("auth", index)
	if err != nil {
		panic("auth configuration error:" + index)
	}
	return e
}
