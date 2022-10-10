package services

import (
	"errors"
	"github.com/lavinas/payly-service/internal/core/domains"
	"github.com/lavinas/payly-service/internal/core/ports"
	"golang.org/x/crypto/bcrypt"
	"reflect"
)

type authenticate struct {
	user   ports.User
	config ports.Config
	jwt    ports.AuthJWT
	log    ports.Log
}

func NewAuthenticate(user ports.User, config ports.Config, jwt ports.AuthJWT, log ports.Log) *authenticate {
	return &authenticate{user: user, config: config, jwt: jwt, log: log}
}

func (a *authenticate) Token(auth domains.AuthIn) (domains.AuthToken, error) {
	if err := a.validate(auth); err != nil {
		a.log.Info(auth.Username + " - " + err.Error())
		return domains.AuthToken{}, err
	}
	if err := a.checkClient(auth); err != nil {
		a.log.Info(auth.Username + " - " + err.Error())
		return domains.AuthToken{}, err
	}
	id, err := a.login(auth)
	if err != nil {
		a.log.Info(auth.Username + " - " + err.Error())
		return domains.AuthToken{}, err
	}
	code, lag, err := a.jwt.Get(auth.Id, auth.Username, id)
	if err != nil {
		a.log.Info(auth.Username + " - " + err.Error())
		return domains.AuthToken{}, errors.New("internal_error: generating token")
	}
	t := domains.AuthToken{
		Code:   code,
		Type:   "Bearer",
		Expire: lag,
	}
	a.log.Info(auth.Username + " - OK")
	return t, nil
}

func (a *authenticate) validate(auth domains.AuthIn) error {
	message := "invalid_request: the request is missing a required parameter, " +
		"includes an invalid parameter value, includes a " +
		"parameter more than once, or is otherwise malformed. Check "
	hasError := false
	values := reflect.ValueOf(auth)
	typesOf := values.Type()
	for i := 0; i < values.NumField(); i++ {
		p := values.Field(i).Interface().(string)
		if p == "" {
			message += typesOf.Field(i).Tag.Get("json") + ", "
			hasError = true
		}
	}
	if hasError {
		message = message[0 : len(message)-2]
		return errors.New(message)
	}
	return nil
}

func (a *authenticate) checkClient(auth domains.AuthIn) error {
	id, err := a.config.GetField("auth", "client_id")
	if err != nil {
		return errors.New("internal_error: client_id config error")
	}
	secret, err := a.config.GetField("auth", "client_secret")
	if err != nil {
		return errors.New("internal_error: client_secret config error")
	}
	if auth.Id != id || auth.Secret != secret {
		return errors.New("invalid_client: Client authentication failed")
	}
	return nil
}

func (a *authenticate) login(auth domains.AuthIn) (int, error) {
	i, pass, err := a.user.GetActive(auth.Username)
	if err != nil {
		return 0, errors.New("invalid_credentials: The user credentials were incorrect")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(pass), []byte(auth.Password)); err != nil {
		return 0, errors.New("invalid_credentials: The user credentials were incorrect")
	}
	return i, nil
}
