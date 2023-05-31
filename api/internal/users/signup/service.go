package signup

import (
	"context"
	"errors"
	"fmt"
	"github.com/software-advice/aries-team-assessment/internal/users"
)

// ErrAlreadyTakenUsername is the error to return if there is already a User with the given username
var ErrAlreadyTakenUsername = errors.New("username already taken")

// ErrFindingUser is the error to return when there is a platform problem finding the user
var ErrFindingUser = errors.New("can' find user")

// ErrMakingUser is the error to return when the user can't be made because a wrong given param (as an invalid password)
var ErrMakingUser = errors.New("can't create user")

// ErrSavingUser is the error to returno when ther is a platform problem persisting the User
var ErrSavingUser = errors.New("can't save user")

// Service is an application service to handle a User signup
type Service struct {
	repository     users.Repository
	findingService users.FindingService
}

// SignUp create a new user with th given username and password.
// If there is already a USer with the given username, it returns a ErrAlreadyTakenUsername
// Other possible errors are ErrFindingUser, ErrMakingUser and ErrSavingUser.
func (s Service) SignUp(ctx context.Context, username string, password []byte) error {
	usr, err := s.findingService.Find(ctx, username)
	if err != nil {
		return fmt.Errorf("%w - %v", ErrFindingUser, err)
	}
	if !usr.IsNotUser() { // if some user was found the user was already taken
		return ErrAlreadyTakenUsername
	}
	usr, err = users.MakeUser(username, password)
	if err != nil {
		return fmt.Errorf("%w - %v", ErrMakingUser, err)
	}
	err = s.repository.Save(ctx, usr)
	if err != nil {
		return fmt.Errorf("%w - %v", ErrSavingUser, err)
	}
	return nil
}

// BuildService returns a Service with the given repository
func BuildService(repository users.Repository) Service {
	return Service{
		repository:     repository,
		findingService: users.BuildFindingService(repository),
	}
}
