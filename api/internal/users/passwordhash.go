package users

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type PasswordHash struct {
	val []byte
}

func (h PasswordHash) Bytes() []byte {
	return h.val
}

func MakePasswordHash(password Password) (PasswordHash, error) {
	hash, err := bcrypt.GenerateFromPassword(password.Bytes(), bcrypt.DefaultCost)
	if err != nil {
		return PasswordHash{}, err
	}
	return PasswordHash{val: hash}, nil
}

func ParsePasswordHash(val []byte) (PasswordHash, error) {
	if len(val) == 0 {
		return PasswordHash{}, errors.New("passowrd hash can't be empty")
	}

	return PasswordHash{val: val}, nil
}
