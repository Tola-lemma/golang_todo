package main

import (
	"log"
	"net/http"

	database "github.com/Tola-lemma/golang_todo/Database"
	migrations "github.com/Tola-lemma/golang_todo/Migrations"
	route "github.com/Tola-lemma/golang_todo/Route"
	"github.com/spf13/viper"
)

func main() {
	database.InitDB()
	migrations.RunMigrations()
	r:=route.InitRoutes()
	port:=viper.GetString("SERVER_PORT")
	if port ==""{
		port="8080"
	}
	log.Printf("Server running on port %s",port)
	http.ListenAndServe(":"+port,r)
}