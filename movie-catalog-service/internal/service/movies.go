package service

import (
	"context"
	"zonos.com/movie_catalog_microservices/internal/models"
)

type MovieStore interface {
	GetMovie(id string) (*models.Movie, error)
	GetMovies() *models.Movies
}

type MovieService struct {
	MovieStore MovieStore
}

func NewMovieService(store MovieStore) MovieService {
	return MovieService{
		MovieStore: store,
	}
}

func (ms *MovieService) GetMovie(ctx context.Context, id string) (*models.Movie, error) {
	movie, err := ms.MovieStore.GetMovie(id)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (ms *MovieService) GetMovies(ctx context.Context) *models.Movies {
	return ms.MovieStore.GetMovies()
}
