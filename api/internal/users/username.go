package users

import "errors"

type Username struct {
	val string
}

func (u Username) String() string {
	return u.val
}

func ParseUsername(val string) (Username, error) {
	if val == "" {
		return Username{}, errors.New("username can't be empty")
	}
	return Username{val: val}, nil
}

func ParseUnsafeUsername(val string) Username {
	return Username{val: val}
}
