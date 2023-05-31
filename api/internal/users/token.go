package users

// Token represents a TokenString with its Claims.
type Token struct {
	claims Claims
	str    TokenString
}

// Claims returns the Calims of the Token.
func (t Token) Claims() Claims {
	return t.claims
}

// TokenString returns the TokenString.
func (t Token) TokenString() TokenString {
	return t.str
}

// TokenString object value used to represent a token string.
type TokenString struct {
	val string
}

// String returns the string representation of the TokenString.
func (t TokenString) String() string {
	return t.val
}

// ParseTokenString parses the given string to a TokenString.
func ParseTokenString(strTkn string) TokenString {
	return TokenString{
		val: strTkn,
	}
}

// A TokenGenerator is a struct that can generate a TokenString.
type TokenGenerator interface {
	// Generate returns a TokenString for the given Claims.
	Generate(claims Claims) (TokenString, error)
}

// A TokenValidator is a struct that can validate if a TokenString is valid and gets its Claims.
type TokenValidator interface {
	// Validate returns the token Claims it the token is valid or error if it's not.
	Validate(token TokenString) (Claims, error)
}
