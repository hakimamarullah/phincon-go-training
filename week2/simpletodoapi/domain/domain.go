package domain

import "github.com/google/uuid"

type User struct {
	Id       string `json:"id"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

func NewUser(username string, email string) *User {
	return &User{Id: uuid.New().String(), Username: username, Email: email}
}

type Todo struct {
	Id          string `json:"id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	UserId      string `json:"uid" validate:"required"`
	IsFinished  bool   `json:"isFinished"`
}

func NewTodo(title string, description string, uid string) *Todo {
	return &Todo{Id: uuid.New().String(), Title: title, Description: description, UserId: uid, IsFinished: false}
}

type ResponseBody struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Count   int         `json:"count"`
	Code    int         `json:"code"`
}

func NewResponseBody(message string, data interface{}, count int, code int) *ResponseBody {
	return &ResponseBody{Message: message, Data: data, Count: count, Code: code}
}
