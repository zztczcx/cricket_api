package db

import (
	"cricket/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Store interface {
        Querier
}

type SQLStore struct {
        *Queries
}

func NewStore(db DBTX) Store {
        return &SQLStore{
                Queries: New(db),
        }
}

func NewDatabase(cfg config.Database) (*sql.DB, error) {
	connPool, err := sql.Open("mysql", cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	connPool.SetConnMaxLifetime(cfg.MaxConnectionLifeTime)
	connPool.SetMaxOpenConns(cfg.MaxOpenConnections)
	connPool.SetMaxIdleConns(cfg.MaxIdleConnections)

	return connPool, nil
}
