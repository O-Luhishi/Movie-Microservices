package main

import (
	"log"
	"zonos.com/movie_catalog_microservices/internal/config"
	"zonos.com/movie_catalog_microservices/internal/db"
	"zonos.com/movie_catalog_microservices/internal/repository"
	"zonos.com/movie_catalog_microservices/internal/service"
	"zonos.com/movie_catalog_microservices/internal/transport/grpc"
)

func Run() error {

	appConf := config.NewAppConfig()
	dbConn, err := db.NewPostgresDB(appConf)
	if err != nil {
		log.Printf("Could Not Connect To The Database")
	}
	movieRepo := repository.NewMovieStore(dbConn)
	movieSvc := service.NewMovieService(movieRepo)
	movieHandler := grpc.NewMovieGrpcServer(&movieSvc, *appConf)

	if err := movieHandler.Serve(); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
