package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type dollars float32

// Catalog of products
type Catalog map[string]dollars

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()

	query := "/search"
	ls, err := search(ctx, query)

	if err != nil {
		log.Fatal(err)
	}

	for prod, price := range ls {
		fmt.Printf("Product: %s, Price: %s\n", prod, price)
	}
}

func search(ctx context.Context, query string) (Catalog, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8000/search?timeout=1s", nil)

	if err != nil {
		log.Fatal(err)
	}
	req = req.WithContext(ctx)

	return httpDo(ctx, req, func(resp *http.Response, err error) (Catalog, error) {
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		var c Catalog
		if err := json.NewDecoder(resp.Body).Decode(&c); err != nil {
			return nil, err
		}
		return c, nil
	})
}

func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) (Catalog, error)) (Catalog, error) {

	type results struct {
		c   Catalog
		err error
	}

	c := make(chan results, 1)

	req = req.WithContext(ctx)

	go func() {
		ls, err := f(http.DefaultClient.Do(req))
		c <- results{ls, err}
	}()

	select {
	case <-ctx.Done():
		<-c
		return nil, ctx.Err()
	case rs := <-c:
		return rs.c, rs.err
	}
}
