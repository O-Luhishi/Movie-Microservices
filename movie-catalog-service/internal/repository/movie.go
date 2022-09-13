package repository

import (
	"github.com/jinzhu/gorm"
	"zonos.com/movie_catalog_microservices/internal/models"
)

type MovieRepository struct {
	db *gorm.DB
}

func NewMovieStore(db *gorm.DB) MovieRepository {
	return MovieRepository{
		db: db,
	}
}

func (m MovieRepository) GetMovie(id string) (*models.Movie, error) {
	if id == "1" {
		return &models.Movie{
			MovieID:     "01",
			MovieName:   "Lord Of The Rings",
			ReleaseDate: "January 1st 2001",
			Description: "Film About A Ring",
		}, nil
	} else if id == "2" {
		return &models.Movie{
			MovieID:     "02",
			MovieName:   "The Matrix",
			ReleaseDate: "February 1st 1999",
			Description: "Film About Simulation",
		}, nil
	} else if id == "3" {
		return &models.Movie{
			MovieID:     "03",
			MovieName:   "GoodFellas",
			ReleaseDate: "January 12st 1998",
			Description: "Film About A Mafia",
		}, nil
	} else {
		return &models.Movie{}, nil
	}
}

func (m MovieRepository) GetMovies() *models.Movies {
	return &models.Movies{Movies: []models.Movie{
		{
			MovieID:     "01",
			MovieName:   "Lord Of The Rings",
			ReleaseDate: "January 1st 2001",
			Description: "Film About A Ring",
		},
		{
			MovieID:     "02",
			MovieName:   "The Matrix",
			ReleaseDate: "February 1st 1999",
			Description: "Film About Simulation",
		},
		{
			MovieID:     "03",
			MovieName:   "GoodFellas",
			ReleaseDate: "January 12st 1998",
			Description: "Film About A Mafia",
		},
	}}
}
