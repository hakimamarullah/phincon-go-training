package main

import (
	"fmt"
	"net/http"
	"simpletodoapi/handler"
	"simpletodoapi/middlewares"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", Users)
	mux.HandleFunc("/todos", Todos)
	mux.HandleFunc("/login", Login)
	mux.HandleFunc("/batch/delete_todo", ClearTodo)

	var serveMux http.Handler = mux
	serveMux = middlewares.Logging(serveMux)
	serveMux = middlewares.Authenticate(serveMux)
	serveMux = middlewares.HandlerAdvice(serveMux)
	server := http.Server{
		Addr:    "localhost:8000",
		Handler: serveMux,
	}

	fmt.Printf("Server is up and running on %s\n", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func Users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handler.AddUser(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
	}
}

func Todos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handler.GetAllTodoByUserId(w, r)
	case http.MethodPost:
		handler.AddTodo(w, r)
	case http.MethodDelete:
		handler.DeleteTodoById(w, r)
	case http.MethodPut:
		handler.UpdateTodo(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handler.Login(w, r)
	}
}

func ClearTodo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		handler.DeleteAllTodoByUser(w, r)
	}
}
