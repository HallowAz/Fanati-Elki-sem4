package main

import (
	"fmt"
	"net/http"

	"fe-sem4/config"
	form_handlers_lib "fe-sem4/internal/handlers/form"
	form_managers_lib "fe-sem4/internal/managers/form"
	"fe-sem4/internal/repository"
)

func main() {

	formRepo := repository.NewFormRepo()
	formManager := form_managers_lib.NewManager(formRepo)
	formHandler := form_handlers_lib.NewFormHandler(formManager)

	// Просто для проверки работоспособности
	// Можно на гориллу переделать, если удобно с ней
	http.HandleFunc("/", formHandler.CreateForm)

	fmt.Printf("Server start at: %s", config.ServerPort)
	err := http.ListenAndServe(config.ServerPort, nil)
	if err != nil {
		fmt.Println(err)
	}
}
