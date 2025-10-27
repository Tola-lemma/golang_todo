package route

import (
	controller "github.com/Tola-lemma/golang_todo/Controller"
	middleware "github.com/Tola-lemma/golang_todo/Middleware"
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router{
	r:=mux.NewRouter()
	r.Use(middleware.Logger)
	api:=r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/todos",controller.GetTodos).Methods("GET")
	api.HandleFunc("/todos",controller.CreateTodo).Methods("POST")
	api.HandleFunc("/todos/{id}", controller.UpdateTodo).Methods("PUT")
	api.HandleFunc("/todos/{id}", controller.DeleteTodo).Methods("DELETE")
	return  r
}