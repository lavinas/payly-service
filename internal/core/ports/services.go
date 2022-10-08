package ports

import (
	"github.com/lavinas/payly-service/internal/core/domains"
)

type AuthService interface {
	Token(domains.AuthIn) (domains.AuthToken, error)
}
