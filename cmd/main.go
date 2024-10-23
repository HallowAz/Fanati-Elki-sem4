package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"fe-sem4/config"
	"fe-sem4/infra"
	problem_handlers_lib "fe-sem4/internal/handlers/problem"
	problem_managers_lib "fe-sem4/internal/managers/problem"
	"fe-sem4/internal/repository"
	"fe-sem4/internal/repository/db"
	"fe-sem4/metrics"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()
	router := mux.NewRouter()

	err := config.InitConfig(config.ConfigEnvFilePath)
	if err != nil {
		log.Fatal(err)
	}

	pool, err := infra.ConnectToDB(ctx)
	if err != nil {
		log.Fatal(err)
	}

	dbTX := db.NewTXCommitter(pool)

	problemRepo := repository.NewProblemRepo(dbTX)
	problemManager := problem_managers_lib.NewManager(problemRepo)
	problemHandler := problem_handlers_lib.NewFormHandler(problemManager, problemRepo)

	problemHandler.RegisterRoutes(router)

	go func() {
		_ = metrics.Listen("127.0.0.1:8082")
	}()

	fmt.Printf("Server start at: %s\n", config.ServerPort)
	err = http.ListenAndServe(config.ServerPort, router)
	if err != nil {
		log.Fatal(err)
	}
}
