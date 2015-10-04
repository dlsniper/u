package main

import (
	"fmt"
	"os/user"
)

func main() {
	fmt.Println(getUserName())
}

func getUserName() string {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	return user.Username
}