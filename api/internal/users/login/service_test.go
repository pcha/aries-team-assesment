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
	"time"
)

func TestService_Login(t *testing.T) {
	validPAsswordHash := []byte("$2a$10$Vq8Tx8eLAFevAULXWtfJXOFFh6eMAMgJ4rQwPett62hO6.6zCJ9eW")
	type useCase struct {
		username string
		password string
		mockUser users.User
		wantErr  error
	}
	cases := map[string]useCase{
		"ok": {
			username: "test",
			password: "asd123",
			mockUser: users.BuildUnsafe("test", validPAsswordHash),
			wantErr:  nil,
		},
		"wrong password": {
			username: "test",
			password: "wrong password",
			mockUser: users.BuildUnsafe("test", validPAsswordHash),
			wantErr:  users.ErrInvalidPassword,
		},
		"wrong username": {
			username: "testt",
			password: "asd123",
			mockUser: users.NotUsr,
			wantErr:  ErrUserNotFound,
		},
	}
	for caseName, tc := range cases {
		t.Run(caseName, func(t *testing.T) {

			repo := new(mockable.UsersRepository)
			tknGen, err := jwt.BuildHS256Manager([]byte("testKey"))
			require.NoError(t, err)
			tknGenSrvc := users.BuildTokenGenerationService(tknGen, time.Second)
			service := BuildService(repo, tknGenSrvc)

			repo.On("Get", mock.Anything, mock.MatchedBy(func(usrName users.Username) bool {
				return usrName.String() == tc.username
			})).
				Return(tc.mockUser, nil)

			tkn, err := service.Login(context.Background(), tc.username, []byte(tc.password))

			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, tkn.TokenString().String())
			}
			repo.AssertExpectations(t)
		})
	}
}
