package login

import (
	"context"
	"fmt"
	"github.com/software-advice/aries-team-assessment/internal/users"
)

type Service struct {
	repository     users.Repository
	tokenGenerator users.TokenGenerator
}

type Token string

func (t Token) String() string {
	return string(t)
}

func (s Service) Login(ctx context.Context, username string, password []byte) (users.Token, error) {
	notTkn := users.Token{}

	user, err := s.repository.Get(ctx, users.ParseUnsafeUsername(username))
	if err != nil {
		return notTkn, err
	}
	err = user.ValidatePassword(password)
	if err != nil {
		return notTkn, err
	}

	claims := users.BuildClaims(user)
	tkn, err := s.tokenGenerator.Generate(claims)
	if err != nil {
		return notTkn, fmt.Errorf("token error: %v", err)
	}
	return tkn, nil
}

func BuildService(repository users.Repository, tknGen users.TokenGenerator) Service {
	return Service{
		repository:     repository,
		tokenGenerator: tknGen,
	}
}
