package main

import (
	"fmt"
	"log"
	"net/http"

	"fe-sem4/config"
	"fe-sem4/db"
	user_handlers_lib "fe-sem4/internal/handlers/user"
	user_managers_lib "fe-sem4/internal/managers/user"
	"fe-sem4/internal/repository"
	_ "github.com/lib/pq"
)

var (
	host   = "83.166.237.142"
	port   = 5432
	user   = "postgres"
	dbname = "fe-sem4"

	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, "password123", dbname)
)

func main() {
	db, err := db.GetPostgres(psqlInfo)
	if err != nil {
		fmt.Println(err, " ", psqlInfo)
		log.Fatalf("cant connect to postgres")
		return
	}
	defer db.Close()

	userRepo := repository.NewUserRepo(db)
	userManager := user_managers_lib.NewManager(userRepo)
	userHandler := user_handlers_lib.NewUserHandler(userManager, userRepo)

	// Просто для проверки работоспособности
	// Можно на гориллу переделать, если удобно с ней
	http.HandleFunc("/signup", userHandler.CreateUser)

	fmt.Printf("Server start at: %s", config.ServerPort)
	err = http.ListenAndServe(config.ServerPort, nil)
	if err != nil {
		fmt.Println(err)
	}
}
