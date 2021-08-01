package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/robfig/cron"

	"github.com/gorilla/mux"
	helloBusiness "github.com/sebastianaldi17/dockerfullstack/server/code/business/hello"
	helloData "github.com/sebastianaldi17/dockerfullstack/server/code/data/hello"
	helloHandler "github.com/sebastianaldi17/dockerfullstack/server/code/handlers/hello"
)

type clients struct {
	helloData     *helloData.Data
	helloBusiness *helloBusiness.Business
	helloHandler  *helloHandler.Handler
}

func main() {
	// initialize db
	connStr := "postgres://root:root@postgres/docker-fullstack-db?sslmode=disable"
	db, err := sql.Open("postgres", connStr) // change _ to db
	if err != nil {
		log.Fatal(err)
	}

	var c clients

	// initialize data layer
	helloData := helloData.New(db)
	c.helloData = helloData

	// initialize business layer
	helloBusiness := helloBusiness.New(helloData)
	c.helloBusiness = helloBusiness

	// initialize handlers
	helloHandler := helloHandler.New(helloBusiness)
	c.helloHandler = helloHandler

	// initialize crons
	initializeCrons(c)

	// initialize http routes
	initializeRoutes(db, c)

}

func initializeRoutes(db *sql.DB, clients clients) {
	router := mux.NewRouter().StrictSlash(false)

	router.HandleFunc("/", clients.helloHandler.Hello).Methods("GET")
	router.HandleFunc("/logs", clients.helloHandler.GetLogs).Methods("POST")

	log.Println("Backend server started on port 8080.")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func initializeCrons(clients clients) {
	crons := cron.New()

	// Logs "Hello Cron" every minute
	// Syntax: https://pkg.go.dev/github.com/robfig/cron
	crons.AddFunc("0 * * * * *", clients.helloHandler.HelloCron)

	crons.Start()
}
