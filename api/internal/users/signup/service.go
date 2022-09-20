package signup

import (
	"context"
	"errors"
	"fmt"
	"github.com/software-advice/aries-team-assessment/internal/users"
)

var ErrMakingUser = errors.New("can't create user")
var ErrSavingUser = errors.New("can't save user")

type Service struct {
	repository users.Repository
}

func (s Service) SignUp(ctx context.Context, username string, password []byte) error {
	usr, err := users.MakeUser(username, password)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrMakingUser, err)
	}
	err = s.repository.Save(ctx, usr)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrSavingUser, err)
	}
	return nil
}

func BuildService(repository users.Repository) Service {
	return Service{repository: repository}
}
