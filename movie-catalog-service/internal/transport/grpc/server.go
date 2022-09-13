package grpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"zonos.com/movie_catalog_microservices/internal/config"
	"zonos.com/movie_catalog_microservices/internal/models"
	"zonos.com/movie_catalog_microservices/pb"
)

type MovieService interface {
	GetMovie(ctx context.Context, id string) (*models.Movie, error)
	GetMovies(ctx context.Context) *models.Movies
}

type MovieServer struct {
	pb.UnimplementedMovieCatalogServiceServer
	MovieService MovieService
	Config       config.Conf
}

func NewMovieGrpcServer(movieService MovieService, conf config.Conf) *MovieServer {
	return &MovieServer{
		MovieService: movieService,
		Config:       conf,
	}
}

func (ms MovieServer) Serve() error {
	lis, err := net.Listen("tcp", ms.Config.Server.Port)
	if err != nil {
		log.Printf("could not listen on port: %s \n", ms.Config.Server.Port)
		return err
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMovieCatalogServiceServer(grpcServer, &ms)

	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("failed to serve: %s\n", err)
		return err
	}
	return nil
}

func (ms *MovieServer) GetMovie(ctx context.Context, req *pb.GetMovieRequest) (*pb.Movie, error) {
	movie, err := ms.MovieService.GetMovie(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Movie not found")
	}
	return &pb.Movie{
		MovieId:     movie.MovieID,
		MovieName:   movie.MovieName,
		ReleaseDate: movie.ReleaseDate,
		Description: movie.Description,
	}, nil
}
func (ms *MovieServer) GetMovieList(ctx context.Context, req *pb.GetMovieListRequest) (*pb.GetMovieListResponse, error) {
	var movies = &models.Movies{}
	movies = ms.MovieService.GetMovies(ctx)
	return movies.MoviesDTO(), nil
}
