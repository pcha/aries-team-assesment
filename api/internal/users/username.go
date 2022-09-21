package users

import "errors"

// Username is the object value to represent a username
type Username struct {
	val string
}

// String return the String representation of the Username
func (u Username) String() string {
	return u.val
}

// ParseUsername tries to parse the given string to a Username.
// If the string is a not valid username it returns error.
func ParseUsername(val string) (Username, error) {
	if val == "" {
		return Username{}, errors.New("username can't be empty")
	}
	return Username{val: val}, nil
}

// ParseUnsafeUsername parses the string to username without performa any validation.
func ParseUnsafeUsername(val string) Username {
	return Username{val: val}
}
