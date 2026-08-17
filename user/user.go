package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName  string
	lastName   string
	birthdate  string
	createdAtt time.Time
}

// Nests the function inside of user struct
func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *User) ClearUserName () {
	u.firstName = ""
	u.lastName = ""
}

func New() (*User, error) {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birthday (MM/DD/YYYY): ")

	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("First name, last name, and birthdate are required.")
	}

	return &User{
		firstName: firstName,
		lastName: lastName,
		birthdate: birthDate,
		createdAtt: time.Now(),
	}, nil
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}