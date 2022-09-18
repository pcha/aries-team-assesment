package users

type Password struct {
	val []byte
}

func (p Password) Bytes() []byte {
	return p.val
}

func ParsePassword(val []byte) (Password, error) {
	// TODO validations
	return Password{val: val}, nil
}
