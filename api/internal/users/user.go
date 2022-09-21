package users

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var ErrInvalidPassword = errors.New("invalid password")

// User is the model to represent a User
type User struct {
	username     Username
	passwordHash PasswordHash
}

// NotUsr is an empty User
var NotUsr = User{}

// IsNotUser returns true if the User is empty
func (u User) IsNotUser() bool {
	return u.username == NotUsr.username &&
		u.passwordHash.Bytes() == nil
}

// Username returns the User Username
func (u User) Username() Username {
	return u.username
}

// PasswordHash returns the User PasswordHash
func (u User) PasswordHash() PasswordHash {
	return u.passwordHash
}

// ValidatePassword if the given password is the corresponding to the User
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

// MakeUser returns a User wit the given username and password if they are valid, if they aren't returns an error.
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

// BuildUnsafe returns a User with the given user and password hash without validate them.
func BuildUnsafe(username string, passwordHash []byte) User {
	return User{
		username:     Username{val: username},
		passwordHash: PasswordHash{val: passwordHash},
	}
}
