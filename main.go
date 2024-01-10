package main

import (
	"fmt"
	"go-todo-2/app/models"
)

func main() {
	// fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.Password="testpass"
	// fmt.Println(u)

	// u.CreateUser()

	// user, _ := models.GetUser(1)
	// fmt.Println(user)

	// user.Name = "test2"
	// user.Email = "test@test.com"
	// user.UpdateUser()

	// user, _ = models.GetUser(1)
	// fmt.Println(user)

	// user.DeleteUser()

	user, _ := models.GetUser(3)

	user.CreateTodo("test todo")

	todos, _ := user.GetTodosByUser()
	for _, v := range todos {
		fmt.Println(v)
	}

	// fmt.Println(user)
	// user.CreateTodo("First todo")

	// todo, _ := models.GetTodo(1)
	// fmt.Println(todo)
	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }
}
