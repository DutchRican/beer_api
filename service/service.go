package service

import (
	"database/sql"
	"fmt"
)

type ConnectionOptions struct {
	Port     string
	Username string
	Password string
	DB_port  string
	DP_Ip    string
}

type DBConnection interface {
	Open(options ConnectionOptions) error
	Close() error
	Instance() *sql.DB
}

type DB struct {
	Db *sql.DB
}

func (d *DB) Open(options ConnectionOptions) error {
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/beers?sslmode=disable",
		options.Username, options.Password, options.DP_Ip, options.DB_port)

	pg, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	pg.Exec(CreateSchema)
	d.Db = pg
	return nil
}

func (d *DB) Close() error {
	return d.Db.Close()
}
