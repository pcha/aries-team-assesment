package users

import "errors"

// ErrEmptyPassword is the error to return if the password is empty
var ErrEmptyPassword = errors.New("password can't be empty")

// Password is the object value used to represent a User password
type Password struct {
	val []byte
}

// Bytes returns the []byte representation of the Password
func (p Password) Bytes() []byte {
	return p.val
}

// ParsePassword tries to parse the given []byte to a Password.
// if the given []byte is empty it returns an ErrEmptyPassword
func ParsePassword(val []byte) (Password, error) {
	if len(val) == 0 {
		return Password{}, ErrEmptyPassword
	}
	return Password{val: val}, nil
}
