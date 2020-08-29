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

	// we want to get search results within 100ms

	req, err := http.NewRequest("GET", "http://localhost:8000/search?timeout=0.1s", nil)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	catalog, err := httpDo(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	for prod, price := range catalog {
		fmt.Printf("Product: %s, Price: %s\n", prod, price)
	}
}

func httpDo(ctx context.Context, req *http.Request) (Catalog, error) {
	type results struct {
		catalog Catalog
		err     error
	}

	ch := make(chan results, 1)

	go func() {
		defer close(ch)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			ch <- results{nil, err}
			return
		}
		defer resp.Body.Close()

		var catalog Catalog
		if err := json.NewDecoder(resp.Body).Decode(&catalog); err != nil {
			ch <- results{nil, err}
			return
		}
		select {
		case <-ctx.Done():
			ch <- results{nil, ctx.Err()}
			return
		case ch <- results{catalog, err}:
		}

	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case result := <-ch:
		return result.catalog, result.err
	}
}
