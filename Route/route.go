package route

import (
	"github.com/gorilla/mux"
    middleware "github.com/Tola-lemma/golang_todo/Middleware"
)

func InitRoutes() *mux.Router{
	r:=mux.NewRouter()
	r.Use(middleware.Logger)
	
	return  r
}