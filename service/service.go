package service

import (
	"fmt"

	"github.com/dutchrican/beer_api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ConnectionOptions struct {
	Host     string
	User     string
	Password string
	Port     string
	Dbname   string
}

type DB struct {
	Db *gorm.DB
}

func (d *DB) Open(options ConnectionOptions) error {
	connStr := fmt.Sprintf("host=db port=%s user=%s password=%s dbname=%s sslmode=disable",
		options.Port, options.User, options.Password, options.Dbname)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	// pg, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	d.Db = db

	db.AutoMigrate(&models.Country{}, &models.Beer{})
	return nil
}

func (d *DB) Close() error {
	return d.Close()
}
