package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

func slowQuery(ctx context.Context) error {
	_, err := db.ExecContext(ctx, "SELECT pg_sleep(5)")
	return err
}

func slowHandler(w http.ResponseWriter, req *http.Request) {
	start := time.Now()
	err := slowQuery(req.Context())
	if err != nil {
		switch {
		case errors.Is(err, context.Canceled):
			log.Printf("Warning: %s\n", err.Error())
		default:
			log.Printf("Error: %s\n", err.Error())
		}
		return
	}
	fmt.Fprintln(w, "OK")
	fmt.Printf("slowHandler took: %v\n", time.Since(start))
}

func main() {
	var err error

	connstr := "host=localhost port=5432 user=alice password=pa$$word  dbname=wonderland sslmode=disable"

	db, err = sql.Open("postgres", connstr)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}
	srv := http.Server{
		Addr:         "localhost:8000",
		WriteTimeout: 2 * time.Second,
		Handler:      http.TimeoutHandler(http.HandlerFunc(slowHandler), 1*time.Second, "Timeout!\n"),
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf("Server failed: %s\n", err)
	}
}
