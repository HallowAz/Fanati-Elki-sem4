package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const PORT = ":8080"

var (
	redisAddr = flag.String("addr", "redis://redis-session:6379/0", "redis addr")

	host     = "test_postgres"
	port     = 5432
	user     = "uliana"
	password = "uliana"
	dbname   = "prinesy-poday"

	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
)

type HelloStruct struct {
	Str string
}

type Result struct {
	Body interface{}
}

type Error struct {
	Err string
}

func GetMap(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	hellostr := HelloStruct{
		Str: "Welcome to VMeste",
	}

	body := map[string]interface{}{
		"hello": hellostr,
	}
	encoder := json.NewEncoder(w)
	err := encoder.Encode(&Result{Body: body})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(&Error{Err: "error while marshalling JSON"})
		return
	}
}

func main() {
	flag.Parse()

	http.HandleFunc("/api/map", GetMap)

	fmt.Println("Server start at port", PORT[1:])

	fmt.Println("To close connection CTRL+C :-)")

	// Spinning up the server.
	err := http.ListenAndServe("127.0.0.1"+PORT, nil)
	if err != nil {
		log.Fatal(err)
	}

}
