package main

import (
	"fmt"
	"example.com/pointers/user"
)

func main() {
	var appUser *user.User
	var err error

	appUser, err = user.New()

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.OutputUserDetails()
	appUser.ClearUserName()
}
