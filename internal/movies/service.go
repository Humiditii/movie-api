package movies

import "context"

type Service interface {
	ListMovies(ctx context.Context) ([]Movie, error)
}

type movieService struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &movieService{repo: r}
}

func (s *movieService) ListMovies(ctx context.Context) ([]Movie, error) {
	return s.repo.GetAll(ctx)
}
