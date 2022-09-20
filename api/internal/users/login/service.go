package login

import (
	"context"
	"fmt"
	"github.com/software-advice/aries-team-assessment/internal/users"
	"time"
)

type Service struct {
	repository     users.Repository
	tokenGenerator users.TokenGenerationService
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

	claims := users.BuildClaims(user.Username(), time.Now())
	tkn, err := s.tokenGenerator.GetNewToken(claims)
	if err != nil {
		return notTkn, fmt.Errorf("token error: %v", err)
	}
	return tkn, nil
}

func BuildService(repository users.Repository, tknGen users.TokenGenerationService) Service {
	return Service{
		repository:     repository,
		tokenGenerator: tknGen,
	}
}
