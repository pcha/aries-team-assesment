package signup

import (
	"context"
	"github.com/software-advice/aries-team-assessment/internal/users"
)

type Service struct {
	repository users.Repository
}

func (s Service) SignUp(ctx context.Context, username string, password []byte) error {
	usr, err := users.MakeUser(username, password)
	if err != nil {
		return err
	}
	return s.repository.Save(ctx, usr)
}

func BuildService(repository users.Repository) Service {
	return Service{repository: repository}
}
