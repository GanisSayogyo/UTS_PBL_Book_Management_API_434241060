package database

import (
	"context"
	"fmt"

	"github.com/GanisSayogyo/UTS_PBL_Book_Management_API_434241060/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectPostgres(cfg config.Config) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}
