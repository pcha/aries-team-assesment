package users_test

import (
	"context"
	"github.com/software-advice/aries-team-assessment/internal/platform/mockable"
	"github.com/software-advice/aries-team-assessment/internal/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindingService_Find(t *testing.T) {
	repo := new(mockable.UsersRepository)
	repo.On("Get", mock.Anything, users.ParseUnsafeUsername("test")).Return(users.NotUsr, nil)
	service := users.BuildFindingService(repo)
	usr, err := service.Find(context.Background(), "test")
	require.NoError(t, err)
	assert.Equal(t, users.NotUsr, usr)
}
