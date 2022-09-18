package users

import (
	"context"
	"errors"
)

var ErrUserNotFound = errors.New("user not found")
var ErrAlreadyTakenUsername = errors.New("username already taken")

type Repository interface {
	Save(ctx context.Context, user User) error
	Get(ctx context.Context, username Username) (User, error)
}
