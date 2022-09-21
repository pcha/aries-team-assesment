package login

import (
	"context"
	"errors"
	"fmt"
	"github.com/software-advice/aries-team-assessment/internal/users"
	"time"
)

// ErrUserNotFound is the error to return if there is not a user with the given username
var ErrUserNotFound = errors.New("user not found")

// Service is an application service to handle a User login
type Service struct {
	repository     users.Repository
	tokenGenerator users.TokenGenerationService
	findingService users.FindingService
}

// Login verifies username and password; if they are correct returns a new token.
// If there is no User with the given username, it returns a ErrUserNotFound error.
// If the User exists but the password is wrong, it returns users.ErrInvalidPassword
func (s Service) Login(ctx context.Context, username string, password []byte) (users.Token, error) {
	notTkn := users.Token{}

	user, err := s.findingService.Find(ctx, username)
	if err != nil {
		return notTkn, err
	}
	if user.IsNotUser() {
		return notTkn, ErrUserNotFound
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
		findingService: users.BuildFindingService(repository),
	}
}
