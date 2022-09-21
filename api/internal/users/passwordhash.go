package users

import (
	"golang.org/x/crypto/bcrypt"
)

// PasswordHash is the object value to represent a hashed password
type PasswordHash struct {
	val []byte
}

// Bytes return the []bite representation of the PasswordHash
func (h PasswordHash) Bytes() []byte {
	return h.val
}

// MakePasswordHash hashes tries to hash the given Password
func MakePasswordHash(password Password) (PasswordHash, error) {
	hash, err := bcrypt.GenerateFromPassword(password.Bytes(), bcrypt.DefaultCost)
	if err != nil {
		return PasswordHash{}, err
	}
	return PasswordHash{val: hash}, nil
}
