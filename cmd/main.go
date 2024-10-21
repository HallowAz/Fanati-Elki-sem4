package main

import (
	"context"
	"fe-sem4/config"
	form_handlers_lib "fe-sem4/internal/handlers/problem"
	form_managers_lib "fe-sem4/internal/managers/problem"
	"fe-sem4/internal/repository"
	"fe-sem4/internal/repository/db"
	"fe-sem4/metrics"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

var (
	host   = "83.166.237.142"
	port   = 5432
	user   = "postgres"
	dbname = "fe-sem4"

	connStr = fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable&connect_timeout=%d",
		"postgres",
		"postgres",       // Замените на имя пользователя
		"password123",    // Замените на пароль
		"83.166.237.142", // Хост базы данных
		"5432",           // Порт базы данных
		"fe-sem4",        // Название базы данных
		5)
)

func main() {
	ctx := context.Background()

	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatalf("error parsing config: %v", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		log.Fatalf("Connect to database failed: %v\n", err)
	}

	fmt.Println("Connection OK!")

	// Проверка соединения с базой данных
	err = pool.Ping(ctx)
	if err != nil {
		log.Fatalf("Ping failed: %v\n", err)
	}

	dbTX := db.NewTXCommitter(pool)

	formRepo := repository.NewProblemRepo(dbTX)
	formManager := form_managers_lib.NewManager(formRepo)
	formHandler := form_handlers_lib.NewFormHandler(formManager, formRepo)

	//// Просто для проверки работоспособности
	//// Можно на гориллу переделать, если удобно с ней
	http.HandleFunc("/problem", formHandler.CreateForm)
	//http.HandleFunc("/problems", formHandler.GetProblems)

	go func() {
		_ = metrics.Listen("127.0.0.1:8082")
	}()

	fmt.Printf("Server start at: %s", config.ServerPort)
	err = http.ListenAndServe(config.ServerPort, nil)
	if err != nil {
		fmt.Println(err)
	}
}
