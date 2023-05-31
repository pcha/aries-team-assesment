package users

import "context"

// FindingService is a domain service to get a User by username
type FindingService struct {
	repository Repository
}

// Find returns the User with the given username if it exists.
// If the User doesn't exist, no error is returned but the value NotUser is returned
func (s FindingService) Find(ctx context.Context, username string) (User, error) {
	return s.repository.Get(ctx, ParseUnsafeUsername(username))
}

// BuildFindingService returns a FindingService with the given repository
func BuildFindingService(repository Repository) FindingService {
	return FindingService{repository: repository}
}
