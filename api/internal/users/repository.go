package users

import (
	"context"
)

// Repository to handle User persistence
type Repository interface {
	// Save persit a User
	Save(ctx context.Context, user User) error
	// Get return a User with the given username.
	//If no User with the given username exists, it returns a NotUsr and no error
	Get(ctx context.Context, username Username) (User, error)
}
