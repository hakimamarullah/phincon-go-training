package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"simpletodoapi/domain"
	"simpletodoapi/repository"
)

func AddUser(r *http.Request) (error, *domain.User) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return errors.New(err.Error()), nil
	}
	defer r.Body.Close()

	exist, _ := repository.FindUserByEmail(user.Email)

	if exist {
		return errors.New("email already used"), nil
	}

	if err := validator.New().Struct(user); err != nil {
		return errors.New("invalid data"), nil
	}

	addedUser := domain.NewUser(user.Username, user.Email)
	repository.AddUser(*addedUser)
	return nil, addedUser
}

func Login(r *http.Request) (error, *http.Cookie) {
	var login = struct {
		Username string `json:"username"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		return err, nil
	}
	defer r.Body.Close()

	exist, user := repository.FindUserByUsername(login.Username)
	if !exist {
		return errors.New("user not found"), nil
	}
	cookie := &http.Cookie{
		Name:     "uid",
		Value:    user.Id,
		MaxAge:   60,
		Path:     "/",
		HttpOnly: true,
	}

	return nil, cookie
}

func AddTodo(r *http.Request) (error, *domain.Todo) {
	var todo domain.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		return errors.New(err.Error()), nil
	}
	defer r.Body.Close()
	cookie, err := r.Cookie("uid")
	if err != nil {
		return errors.New("unauthenticated"), nil
	}
	uid := cookie.Value
	todo.UserId = uid
	if err := validator.New().Struct(todo); err != nil {
		return errors.New("invalid data"), nil
	}

	addedTodo := domain.NewTodo(todo.Title, todo.Description, todo.UserId)
	repository.AddTodo(*addedTodo)
	return nil, addedTodo
}
func GetAllTodoByUserId(r *http.Request) []*domain.Todo {
	cookie, _ := r.Cookie("uid")
	return repository.GetAllTodosByUserId(cookie.Value)
}

func UpdateTodo(r *http.Request) error {
	var todo domain.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		return errors.New("update failed")
	}

	exist, oldTodo := repository.FindTodoById(todo.Id)
	if !exist {
		return errors.New("todo not found")
	}

	if len(todo.Description) != 0 && len(todo.Title) != 0 {
		oldTodo.Description = todo.Description
		oldTodo.Title = todo.Title
	}
	fmt.Println(todo.IsFinished)
	oldTodo.IsFinished = todo.IsFinished
	return nil
}

func DeleteTodoById(r *http.Request) bool {
	query := r.URL.Query()
	id := query.Get("id")
	return repository.DeleteTodoById(id)
}

func DeleteAllTodoByUserId(r *http.Request) (bool, int) {
	cookie, _ := r.Cookie("uid")
	uid := cookie.Value
	return repository.DeleteAllTodoByUserId(uid)
}
