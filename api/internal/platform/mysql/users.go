package mysql

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/software-advice/aries-team-assessment/internal/users"
	"time"
)

// User is a DTO to represent a users.User in de DB
type User struct {
	ID           int64     `db:"id"`
	Username     string    `db:"username"`
	PasswordHash []byte    `db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
}

// UsersRepository is a users.Repository implementation to manage sql persistence
type UsersRepository struct {
	db *sqlx.DB
}

// NewUsersRepository returns a new *UserRepository for the given db object
func NewUsersRepository(db *sqlx.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

// Save add the given user to the `users` table
func (r UsersRepository) Save(ctx context.Context, user users.User) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO `users` (`username`, `password_hash`) VALUES (?, ?)", user.Username().String(), user.PasswordHash().Bytes())
	return err
}

// Get return the user with the given username.
// if the user doesn't exist no error is returned, but the user will be users.NotUsr
func (r UsersRepository) Get(ctx context.Context, username users.Username) (users.User, error) {
	var qryResult []User
	err := r.db.SelectContext(ctx, &qryResult, "SELECT * FROM `users` WHERE `username`=?", username.String())
	if err != nil {
		return users.User{}, fmt.Errorf("mysql couldn't select user: %v", err)
	}
	if len(qryResult) == 0 {
		return users.NotUsr, nil
	}
	return users.BuildUnsafe(qryResult[0].Username, qryResult[0].PasswordHash), nil
}
