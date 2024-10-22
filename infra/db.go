package infra

import (
	"context"
	"log"

	"fe-sem4/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectToDB(ctx context.Context) (*pgxpool.Pool, error) {
	poolConfig, err := pgxpool.ParseConfig(config.DBConnString)
	if err != nil {
		log.Println("Error parsing database connection string")
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		log.Printf("Connect to database failed: %v\n", err)
		return nil, err
	}

	// Проверка соединения с базой данных
	err = pool.Ping(ctx)
	if err != nil {
		log.Printf("Ping failed: %v\n", err)
		return nil, err
	}

	return pool, nil
}
