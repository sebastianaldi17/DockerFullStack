package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
	helloBusiness "github.com/sebastianaldi17/dockerfullstack/server/code/business/hello"
	helloData "github.com/sebastianaldi17/dockerfullstack/server/code/data/hello"
	helloHandler "github.com/sebastianaldi17/dockerfullstack/server/code/handlers/hello"
)

func main() {
	// Init db
	connStr := "postgres://root:root@postgres/docker-fullstack-db?sslmode=disable"
	db, err := sql.Open("postgres", connStr) // change _ to db
	if err != nil {
		log.Fatal(err)
	}

	handlerFuncs(db)
}

func handlerFuncs(db *sql.DB) {
	// initialize data layer
	helloData := helloData.New(db)

	// initialize business layer
	helloBusiness := helloBusiness.New(helloData)

	// initialize handlers
	helloHandler := helloHandler.New(helloBusiness)

	// initialize routes
	router := mux.NewRouter().StrictSlash(false)

	router.HandleFunc("/", helloHandler.Hello).Methods("GET")
	router.HandleFunc("/logs", helloHandler.GetLogs).Methods("POST")

	log.Println("Backend server started on port 8080.")
	log.Fatal(http.ListenAndServe(":8080", router))
}
