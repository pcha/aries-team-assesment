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
	// hash := []byte("$2a$10$Vq8Tx8eLAFevAULXWtfJXOFFh6eMAMgJ4rQwPett62hO6.6zCJ9eW")

	repository.On("Save", mock.Anything, mock.MatchedBy(func(usr users.User) bool {
		err := usr.ValidatePassword(pass)
		return err != nil &&
			usr.Username().String() == username
	}))
	err := service.SignUp(context.Background(), username, pass)
	require.NoError(t, err)
	repository.AssertExpectations(t)
}
