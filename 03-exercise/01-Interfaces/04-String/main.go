package main

import "fmt"

type user struct {
	name  string
	email string
}

// TODO: Implement custom formating for user struct values.

func main() {
	u := user{
		name:  "John Doe",
		email: "johndoe@example.com",
	}
	fmt.Println(u)
}
