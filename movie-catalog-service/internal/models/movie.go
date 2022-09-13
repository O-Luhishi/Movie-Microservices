package models

import "zonos.com/movie_catalog_microservices/pb"

type Movie struct {
	MovieID     string
	MovieName   string
	ReleaseDate string
	Description string
}

type Movies struct {
	Movies []Movie
}

func (ms *Movies) MoviesDTO() *pb.GetMovieListResponse {
	listOfMovies := make([]*pb.Movie, len(ms.Movies))
	for i := range ms.Movies {
		listOfMovies[i] = &pb.Movie{
			MovieId:     ms.Movies[i].MovieID,
			MovieName:   ms.Movies[i].MovieName,
			ReleaseDate: ms.Movies[i].ReleaseDate,
			Description: ms.Movies[i].Description,
		}
	}
	return &pb.GetMovieListResponse{Movies: listOfMovies}
}
