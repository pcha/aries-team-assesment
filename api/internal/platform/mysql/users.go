package mysql

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/software-advice/aries-team-assessment/internal/users"
	"time"
)

type User struct {
	ID           int64     `db:"id"`
	Username     string    `db:"username"`
	PasswordHash []byte    `db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
}

type UsersRepository struct {
	db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

func (r UsersRepository) Save(ctx context.Context, user users.User) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO `users` (`username`, `password_hash`) VALUES (?, ?)", user.Username().String(), user.PasswordHash().Bytes())
	return err
}

func (r UsersRepository) Get(ctx context.Context, username users.Username) (users.User, error) {
	var qryResult []User
	err := r.db.SelectContext(ctx, &qryResult, "SELECT * FROM `users` WHERE `username`=?", username.String())
	if err != nil {
		return users.User{}, fmt.Errorf("mysql couldn't select user: %v", err)
	}
	if len(qryResult) == 0 {
		return users.NotUsr, users.ErrUserNotFound
	}
	return users.BuildFrom(qryResult[0].Username, qryResult[0].PasswordHash), nil
}
