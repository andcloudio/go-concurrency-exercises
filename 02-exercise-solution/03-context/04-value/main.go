package main

import (
	"context"
	"fmt"
)

type userIDKey string
type database map[string]bool

var db database = database{
	"jane": true,
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	processRequest(ctx, "jane")
}

func processRequest(ctx context.Context, userid string) {
	// send userID information to checkMemberShipStatus for
	// database lookup.
	vctx := context.WithValue(ctx,
		userIDKey("userIDKey"),
		userid)

	ch := checkMemberShipStatus(vctx)
	status := <-ch
	fmt.Printf("membership status of userid : %s : %v\n", userid, status)
}

// checkMemberShipStatus - takes context as input.
// extracts the user id information from context.
// spins a goroutine to do database lookup
// sends the result on the returned channel.
func checkMemberShipStatus(ctx context.Context) <-chan bool {
	ch := make(chan bool)
	go func() {
		defer close(ch)
		// do some database lookup
		userid := ctx.Value(userIDKey("userIDKey")).(string)
		status := db[userid]
		ch <- status
	}()
	return ch
}
