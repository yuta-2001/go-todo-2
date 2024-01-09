package main

import (
	"fmt"
	"go-todo-2/app/models"
)

func main() {
	fmt.Println(models.Db)

	u := &models.User{}
	u.Name = "test"
	u.Email = "test@example.com"
	u.Password="testpass"
	fmt.Println(u)

	u.CreateUser()
}