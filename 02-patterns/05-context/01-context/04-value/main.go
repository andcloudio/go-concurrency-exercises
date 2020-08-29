package main

import (
	"context"
	"fmt"
)

type userIDType string
type database map[string]bool

var db database = database{
	"jane": true,
}

func main() {

	// dbLookup does database lookup in a separate goroutine and
	// sends result on the returned channel.
	dbLookup := func(ctx context.Context) <-chan bool {
		ch := make(chan bool)
		go func() {
			defer close(ch)
			// do some database lookup
			userid := ctx.Value(userIDType("userIDKey")).(string)
			status := db[userid]
			ch <- status
		}()
		return ch
	}

	ctx := context.WithValue(context.Background(), userIDType("userIDKey"), "jane")

	ch := dbLookup(ctx)

	status := <-ch

	fmt.Printf("membership status of userid \"jane\": %v\n", status)
}
