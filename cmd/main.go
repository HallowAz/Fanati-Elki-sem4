package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gomodule/redigo/redis"
	_ "github.com/lib/pq"

	"fe-sem4/config"
	"fe-sem4/db"
	session_handlers_lib "fe-sem4/internal/handlers/session"
	user_handlers_lib "fe-sem4/internal/handlers/user"
	session_managers_lib "fe-sem4/internal/managers/session"
	user_managers_lib "fe-sem4/internal/managers/user"
	"fe-sem4/internal/repository"
)

var (
	host   = "83.166.237.142"
	port   = 5432
	user   = "postgres"
	dbname = "fe-sem4"

	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, "password123", dbname)

	redisAddr = flag.String("addr", "redis://127.0.0.1:6379/0", "redis addr")
)

func main() {
	db, err := db.GetPostgres(psqlInfo)
	if err != nil {
		fmt.Println(err, " ", psqlInfo)
		log.Fatalf("cant connect to postgres")
		return
	}
	defer db.Close()

	redisConn, err := redis.DialURL(*redisAddr)
	if err != nil {
		log.Fatal("can`t connect to redis", err)
	}

	userRepo := repository.NewUserRepo(db)
	userManager := user_managers_lib.NewManager(userRepo)
	userHandler := user_handlers_lib.NewUserHandler(userManager)

	sessionRepo := repository.NewSessionRepo(redisConn)
	sessionManager := session_managers_lib.NewManager(sessionRepo, userRepo)
	sessionHandler := session_handlers_lib.NewSessionHandler(sessionManager)

	// Просто для проверки работоспособности
	// Можно на гориллу переделать, если удобно с ней
	http.HandleFunc("/signup", userHandler.SignUp)
	http.HandleFunc("/login", sessionHandler.Login)
	http.HandleFunc("/logout", sessionHandler.Logout)
	http.HandleFunc("/auth", sessionHandler.Auth)

	fmt.Printf("Server start at: %s", config.ServerPort)
	err = http.ListenAndServe(config.ServerPort, nil)
	if err != nil {
		fmt.Println(err)
	}
}
