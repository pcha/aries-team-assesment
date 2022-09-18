package mockable

import (
	"context"
	"github.com/software-advice/aries-team-assessment/internal/users"
	"github.com/stretchr/testify/mock"
)

type UsersRepository struct {
	mock.Mock
}

func (r UsersRepository) Save(ctx context.Context, user users.User) error {
	args := r.Called(ctx, user)
	return args.Error(0)
}

func (r UsersRepository) Get(ctx context.Context, username users.Username) (users.User, error) {
	args := r.Called(ctx, username)
	return args.Get(0).(users.User), args.Error(1)
}
