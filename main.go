package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	args := os.Args
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	// Current User
	fmt.Printf("My tool! Hi %s! (id: %s), Args: %s\n", user.Name, user.Uid, args)
	fmt.Println("Username: " + user.Username)
	fmt.Println("Home Dir: " + user.HomeDir)
}
