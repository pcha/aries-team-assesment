package login

import (
	"context"
	"github.com/software-advice/aries-team-assessment/internal/platform/jwt"
	"github.com/software-advice/aries-team-assessment/internal/platform/mockable"
	"github.com/software-advice/aries-team-assessment/internal/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestService_Login(t *testing.T) {
	repo := new(mockable.UsersRepository)
	tknGen, err := jwt.BuildHS256Manager([]byte("testKey"))
	require.NoError(t, err)
	tknGenSrvc := users.BuildTokenGenerationService(tknGen)
	service := BuildService(repo, tknGenSrvc)

	username := "test"
	pass := []byte("asd123")
	hash := []byte("$2a$10$Vq8Tx8eLAFevAULXWtfJXOFFh6eMAMgJ4rQwPett62hO6.6zCJ9eW")

	repo.On("Get", mock.Anything, mock.MatchedBy(func(usrName users.Username) bool {
		return usrName.String() == username
	})).
		Return(users.BuildFrom(username, hash), nil)

	tkn, err := service.Login(context.Background(), username, pass)
	require.NoError(t, err)
	assert.NotEmpty(t, tkn.TokenString().String())
	repo.AssertExpectations(t)
}
