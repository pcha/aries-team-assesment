package users

type Token struct {
	claims Claims
	str    string
}

func (t Token) String() string {
	return t.str
}

func ParseToken(strTkn string) Token {
	return Token{
		str: strTkn,
	}
}

type TokenGenerator interface {
	Generate(claims Claims) (Token, error)
}

type TokenValidator interface {
	Validate(token Token) (Claims, error)
}
