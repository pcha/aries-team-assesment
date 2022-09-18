package mysql

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/software-advice/aries-team-assessment/internal/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestUsersRepository_Get(t *testing.T) {
	db, mockDB, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewUsersRepository(sqlx.NewDb(db, "mysql"))
	username := "test"
	hash := "$2a$10$Vq8Tx8eLAFevAULXWtfJXOFFh6eMAMgJ4rQwPett62hO6.6zCJ9eW"
	mockDB.ExpectQuery("SELECT *").
		WithArgs(username).
		WillReturnRows(sqlmock.
			NewRows([]string{`id`, `username`, `password_hash`, `created_at`}).
			AddRow(1, username, hash, time.Now()))
	usr, err := repo.Get(context.Background(), users.ParseUnsafeUsername(username))
	require.NoError(t, err)
	assert.Equal(t, username, usr.Username().String())
	assert.Equal(t, []byte(hash), usr.PasswordHash().Bytes())
}
