package main

import (
	"fmt"
	"log"
	"net/http"

	"fe-sem4/config"
	"fe-sem4/db"
	form_handlers_lib "fe-sem4/internal/handlers/form"
	form_managers_lib "fe-sem4/internal/managers/form"
	"fe-sem4/internal/repository"
)

var (
	host   = "test_postgres"
	port   = 5432
	user   = "uliana"
	dbname = "gorod"

	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, "uliana", dbname)
)

func main() {
	db, err := db.GetPostgres(psqlInfo)
	if err != nil {
		fmt.Println(err, " ", psqlInfo)
		log.Fatalf("cant connect to postgres")
		return
	}
	defer db.Close()

	formRepo := repository.NewProblemRepo(db)
	formManager := form_managers_lib.NewManager(formRepo)
	formHandler := form_handlers_lib.NewFormHandler(formManager, formRepo)

	// Просто для проверки работоспособности
	// Можно на гориллу переделать, если удобно с ней
	http.HandleFunc("/newproblem", formHandler.CreateForm)
	http.HandleFunc("/problems", formHandler.CreateForm)

	fmt.Printf("Server start at: %s", config.ServerPort)
	err = http.ListenAndServe(config.ServerPort, nil)
	if err != nil {
		fmt.Println(err)
	}
}
