package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type dollars float32

// Catalog of products
type Catalog map[string]dollars

var prodCatalog = Catalog{
	"shoe": 10.00,
	"sock": 2.50,
	"book": 8.25,
	"pen":  1.65,
}

// Create db
func Create(ctx context.Context) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./file.db")
	if err != nil {
		return nil, err
	}

	stmt, err := db.PrepareContext(ctx, "CREATE TABLE IF NOT EXISTS catalog (id INTEGER PRIMARY KEY, product TEXT, price FLOAT)")

	_, err = stmt.ExecContext(ctx)
	if err != nil {
		return nil, err
	}
	stmt, err = db.PrepareContext(ctx, "INSERT INTO catalog(product, price) values(?, ?)")
	if err != nil {
		return nil, err
	}

	for prod, price := range prodCatalog {
		_, err = stmt.ExecContext(ctx, prod, price)
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}

// Query db
func Query(ctx context.Context, db *sql.DB) (Catalog, error) {
	rows, err := db.QueryContext(ctx, "SELECT product, price FROM catalog")
	if err != nil {
		return nil, err
	}
	catalog := make(Catalog)
	var product string
	var price dollars
	for rows.Next() {
		rows.Scan(&product, &price)
		catalog[product] = price
	}
	return catalog, nil
}
