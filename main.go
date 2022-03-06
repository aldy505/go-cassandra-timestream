package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gocql/gocql"
)

func main() {
	dbURL, ok := os.LookupEnv("CASSANDRA_URL")
	if !ok {
		dbURL = "localhost"
	}

	cluster := gocql.NewCluster(dbURL)
	cluster.Keyspace = "timestream"
	cluster.Timeout = time.Second * 10

	session, err := cluster.CreateSession()
	if err != nil {
		log.Printf("Error creating session: %v", err)
		return
	}
	defer session.Close()

	deps := &Deps{
		DB: session,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	err = deps.Migrate(ctx)
	if err != nil {
		log.Printf("Error migrating: %v", err)
		return
	}

	router := http.NewServeMux()

	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			deps.InsertHandler(rw, r)
			return
		case http.MethodGet:
			deps.GetHandler(rw, r)
			return
		}
	})

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Println("Starting server on port 8080")
	err = server.ListenAndServe()
	if err != nil {
		log.Printf("Error starting server: %v", err)
	}
}
