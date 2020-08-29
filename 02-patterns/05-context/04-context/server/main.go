// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 195.

// Http4 is an e-commerce server that registers the /list and /price
// endpoint by calling http.HandleFunc.
package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/andcloudio/go-concurrency-exercises/02-patterns/05-context/04-context/server/database"
)

type catalogDB struct {
	db *sql.DB
}

func main() {
	var dbinstance catalogDB
	var err error

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dbinstance.db, err = database.Create(ctx)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/search", dbinstance.handleFunc)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func (db catalogDB) handleFunc(w http.ResponseWriter, req *http.Request) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	timeout, err := time.ParseDuration(req.FormValue("timeout"))
	fmt.Println(timeout)
	if err == nil {
		ctx, cancel = context.WithTimeout(req.Context(), timeout)
	} else {
		ctx, cancel = context.WithCancel(req.Context())
	}
	defer cancel()

	catalog, err := database.Query(ctx, db.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	data, err := json.Marshal(catalog)
	if err != nil {
		log.Fatalf("Json marshalling failed: %v", err)
	}
	fmt.Printf("%s\n", data)

	fmt.Fprintf(w, "%s", data)
}
