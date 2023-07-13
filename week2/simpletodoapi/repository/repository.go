package repository

import (
	"simpletodoapi/domain"
	"strings"
)

var users = make([]domain.User, 0)
var todos = make([]domain.Todo, 0)

func AddUser(user domain.User) {
	users = append(users, user)
}

func FindUserById(id string) (bool, *domain.User) {
	for i, user := range users {
		if strings.EqualFold(user.Id, id) {
			return true, &users[i]
		}
	}
	return false, nil
}

func FindUserByUsername(username string) (bool, *domain.User) {
	for i, user := range users {
		if strings.EqualFold(user.Username, username) {
			return true, &users[i]
		}
	}
	return false, nil
}

func FindUserByEmail(email string) (bool, *domain.User) {
	for i, user := range users {
		if strings.EqualFold(user.Email, email) {
			return true, &users[i]
		}
	}
	return false, nil
}

func AddTodo(todo domain.Todo) {
	todos = append(todos, todo)
}

func FindTodoById(id string) (bool, *domain.Todo) {
	for i, todo := range todos {
		if strings.EqualFold(todo.Id, id) {
			return true, &todos[i]
		}
	}
	return false, nil
}

func GetAllTodosByUserId(uid string) []*domain.Todo {
	var results = make([]*domain.Todo, 0)
	for i, todo := range todos {
		if strings.EqualFold(todo.UserId, uid) {
			results = append(results, &todos[i])
		}
	}
	return results
}

func UpdateTodo(newTodo domain.Todo) bool {
	for i, todo := range todos {
		if strings.EqualFold(todo.Id, newTodo.Id) {
			todos = append(todos[:i], todos[i+1:]...)
			todos = append(todos, todo)
			return true
		}
	}
	return false
}

func DeleteTodoById(id string) bool {
	for i, todo := range todos {
		if strings.EqualFold(todo.Id, id) {
			todos = append(todos[:i], todos[i+1:]...)
			return true
		}
	}
	return false
}

func DeleteAllTodoByUserId(uid string) (bool, int) {
	counter := 0
	for i := len(todos) - 1; i >= 0; i-- {
		if strings.EqualFold(todos[i].UserId, uid) {
			todos = append(todos[:i], todos[i+1:]...)
			counter++
		}
	}
	return true, counter
}
