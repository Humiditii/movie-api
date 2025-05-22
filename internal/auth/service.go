package auth

import "context"

type authService struct {
	repo Repository
}

type Service interface {
	SignupUser(ctx context.Context) (interface{}, error)
}

func NewService(r Repository) Service {
	return &authService{repo: r}
}

func( s *authService) SignupUser(ctx context.Context) (interface{}, error) {

	return nil, nil
}
