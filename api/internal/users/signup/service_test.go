package signup

import (
	"context"
	"github.com/software-advice/aries-team-assessment/internal/platform/mockable"
	"github.com/software-advice/aries-team-assessment/internal/users"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestService_SignUp(t *testing.T) {
	repository := new(mockable.UsersRepository)
	service := BuildService(repository)

	username := "test"
	pass := []byte("asd123")

	repository.
		On("Get", mock.Anything, users.ParseUnsafeUsername(username)).
		Return(users.NotUsr, nil)
	repository.
		On("Save", mock.Anything, mock.MatchedBy(func(usr users.User) bool {
			err := usr.ValidatePassword(pass)
			return err == nil &&
				usr.Username().String() == username
		})).Return(nil)
	err := service.SignUp(context.Background(), username, pass)
	require.NoError(t, err)
	repository.AssertExpectations(t)
}
