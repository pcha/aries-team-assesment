package mockable

import (
	"github.com/software-advice/aries-team-assessment/internal/users"
	"github.com/stretchr/testify/mock"
)

type TokenGenerator struct {
	mock.Mock
}

func (t *TokenGenerator) Generate(claims users.Claims) (users.TokenString, error) {
	args := t.Called(claims)
	return args.Get(0).(users.TokenString), args.Error(1)
}
