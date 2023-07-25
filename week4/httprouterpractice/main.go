package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

func main() {
	router := httprouter.New()
	router.ServeFiles("/file/*filepath", http.Dir("./files"))
	router.GET("/body", getRequestBody)
	router.GET("/cookie", cookie)

	server := http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: router,
	}

	fmt.Printf("Server is running on %s", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func getRequestBody(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var tmp = struct {
		Name string `json:"name"`
		Desc string `json:"desc"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&tmp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tmp)

}

func cookie(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	cookie := &http.Cookie{
		Name:   "path_variable",
		Value:  ps.ByName("cookie"),
		Path:   "/",
		MaxAge: int(1 * time.Minute),
	}

	http.SetCookie(w, cookie)
	w.Header().Set("X-Message", ps.ByName("cookie"))
	fmt.Fprintf(w, "Cookie set with key path_variable and value %s", ps.ByName("cookie"))
}
