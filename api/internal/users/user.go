package users

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var ErrInvalidPassword = errors.New("invalid password")

type User struct {
	username     Username
	passwordHash PasswordHash
}

var NotUsr = User{}

func (u User) Username() Username {
	return u.username
}

func (u User) PasswordHash() PasswordHash {
	return u.passwordHash
}

func (u User) ValidatePassword(password []byte) error {
	err := bcrypt.CompareHashAndPassword(u.passwordHash.Bytes(), password)
	if err != nil {
		return ErrInvalidPassword
	}
	return nil
}

func (u User) GetClaims() Claims {
	return Claims{
		username:  u.username,
		expiresAt: time.Now().Add(time.Hour),
	}
}

func MakeUser(username string, password []byte) (User, error) {
	usrName, err := ParseUsername(username)
	if err != nil {
		return NotUsr, err
	}
	pass, err := ParsePassword(password)
	if err != nil {
		return NotUsr, err
	}
	hash, err := MakePasswordHash(pass)
	if err != nil {
		return NotUsr, err
	}

	return User{
		username:     usrName,
		passwordHash: hash,
	}, nil
}

func BuildFrom(username string, passwordHash []byte) User {
	return User{
		username:     Username{val: username},
		passwordHash: PasswordHash{val: passwordHash},
	}
}
