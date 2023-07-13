package handler

import (
	"net/http"
	"simpletodoapi/domain"
	"simpletodoapi/helpers"
	"simpletodoapi/services"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	err, addedUser := services.AddUser(r)
	if err != nil {
		helpers.ResponseJSON(w, domain.ResponseBody{Message: err.Error(), Code: http.StatusBadRequest})
		return
	}

	helpers.ResponseJSON(w, domain.ResponseBody{Data: addedUser, Code: http.StatusCreated})
}

func Login(w http.ResponseWriter, r *http.Request) {
	err, cookie := services.Login(r)
	if err != nil {
		helpers.ResponseJSON(w, domain.ResponseBody{Message: err.Error(), Code: http.StatusUnauthorized})
		return
	}

	http.SetCookie(w, cookie)
	helpers.ResponseJSON(w, domain.ResponseBody{Message: "Login Success", Data: cookie.Value})
}

func GetAllTodoByUserId(w http.ResponseWriter, r *http.Request) {
	results := services.GetAllTodoByUserId(r)
	helpers.ResponseJSON(w, domain.ResponseBody{Data: results, Count: len(results)})
}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	err, addedTodo := services.AddTodo(r)
	if err != nil {
		helpers.ResponseJSON(w, domain.ResponseBody{Message: err.Error(), Code: http.StatusBadRequest})
		return
	}

	helpers.ResponseJSON(w, domain.ResponseBody{Data: addedTodo, Code: http.StatusCreated})
}

func DeleteTodoById(w http.ResponseWriter, r *http.Request) {
	success := services.DeleteTodoById(r)
	if !success {
		helpers.ResponseJSON(w, domain.ResponseBody{Message: "No Deletion"})
		return
	}

	helpers.ResponseJSON(w, domain.ResponseBody{})
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	if err := services.UpdateTodo(r); err != nil {
		helpers.ResponseJSON(w, domain.ResponseBody{Message: err.Error()})
		return
	}

	helpers.ResponseJSON(w, domain.ResponseBody{Message: "todo updated"})
}

func DeleteAllTodoByUser(w http.ResponseWriter, r *http.Request) {
	success, size := services.DeleteAllTodoByUserId(r)
	if !success {
		helpers.ResponseJSON(w, domain.ResponseBody{Message: "No Deletion", Count: size})
		return
	}

	helpers.ResponseJSON(w, domain.ResponseBody{Message: "All todo deleted", Count: size})
}
