package service

import (
	"database/sql"
	"fmt"
)

type ConnectionOptions struct {
	Host     string
	User     string
	Password string
	Port     string
	Dbname   string
}

type DB struct {
	Db *sql.DB
}

func (d *DB) Open(options ConnectionOptions) error {
	connStr := fmt.Sprintf("host=db port=%s user=%s password=%s dbname=%s sslmode=disable",
		options.Port, options.User, options.Password, options.Dbname)

	pg, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	_, err = pg.Exec(CreateSchema)
	if err != nil {
		return err
	}
	d.Db = pg
	return nil
}

func (d *DB) Close() error {
	return d.Db.Close()
}
