package main

import "fmt"

type user struct {
	name  string
	email string
}

// TODO: Implement custom formating for user struct values.

func (user user) String() string {
	return fmt.Sprintf("users name: %s and email: %s ", user.name, user.email)
}

func main() {
	u := user{
		name:  "John Doe",
		email: "johndoe@example.com",
	}
	fmt.Println(u)
}
