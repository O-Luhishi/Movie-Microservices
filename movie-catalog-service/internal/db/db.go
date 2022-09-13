package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"zonos.com/movie_catalog_microservices/internal/config"
)

func NewPostgresDB(conf *config.Conf) (*gorm.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.Db.Host, conf.Db.Port, conf.Db.Username, conf.Db.Password, conf.Db.DbName)
	return gorm.Open("postgres", connStr)
}
