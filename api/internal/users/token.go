package users

type Token struct {
	claims Claims
	str    TokenString
}

func (t Token) Claims() Claims {
	return t.claims
}

func (t Token) TokenString() TokenString {
	return t.str
}

type TokenString struct {
	val string
}

func (t TokenString) String() string {
	return t.val
}

func ParseTokenString(strTkn string) TokenString {
	return TokenString{
		val: strTkn,
	}
}

type TokenGenerator interface {
	Generate(claims Claims) (TokenString, error)
}

type TokenValidator interface {
	Validate(token TokenString) (Claims, error)
}
