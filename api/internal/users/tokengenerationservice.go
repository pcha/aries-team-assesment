package users

import (
	"errors"
	"fmt"
	"time"
)

// ErrGeneratingToken is an error to return if the TokenGenerationService fails to generate a Token
var ErrGeneratingToken = errors.New("error generating token")

// TokenGenerationService is a domain service that generate new Tokens
type TokenGenerationService struct {
	generator      TokenGenerator
	tokensDuration time.Duration
}

// GetNewToken take the given claims and generate a new token updating the expiration
func (s TokenGenerationService) GetNewToken(claims Claims) (Token, error) {
	claims.expiresAt = time.Now().Add(s.tokensDuration)
	token, err := s.generator.Generate(claims)
	if err != nil {
		return Token{}, fmt.Errorf("%w: %v", ErrGeneratingToken, err)
	}
	return Token{
		claims: claims,
		str:    token,
	}, nil
}

// BuildTokenGenerationService returns a TokenGenerationService with the given TokenGenerator and new tokens duration.
func BuildTokenGenerationService(generator TokenGenerator, tokensDuration time.Duration) TokenGenerationService {
	return TokenGenerationService{generator: generator, tokensDuration: tokensDuration}
}
