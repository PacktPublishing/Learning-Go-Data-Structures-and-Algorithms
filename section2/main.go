package main

import (
	"errors"
	"fmt"
)

type User struct {
	Name, Role, Email string
	Age               int
}

func (u User) Salary() (int, error) {
	if u.Role == "" {
		return 0, errors.New("I can't handle empty Roles")
	}

	switch u.Role {
	case "Developer":
		return 100, nil
	case "Architect":
		return 200, nil
	}
	return 0, errors.New(
		fmt.Sprintf("I'm not able to handle the '%s' role", u.Role),
	)
}

func main() {
	user := User{Name: "Gustavo", Role: "Unknown"}
	if salary, err := user.Salary(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Salary:", salary)
	}
}
