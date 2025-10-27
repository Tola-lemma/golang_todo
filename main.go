package main

import (
	"log"
	"net/http"

	config "github.com/Tola-lemma/golang_todo/Config"
	database "github.com/Tola-lemma/golang_todo/Database"
	migrations "github.com/Tola-lemma/golang_todo/Migrations"
	route "github.com/Tola-lemma/golang_todo/Route"
)

func main() {
	config.LoadConfig()
	database.InitDB()
	migrations.RunMigrations()
	r:=route.InitRoutes()
	cfg :=config.AppConfig
	port:=cfg.ServerPort
	serverAddr:=cfg.ServerAddress
	if port ==""{
		port="8080"
	}
	log.Printf("Server running on port %s",port)
	http.ListenAndServe(serverAddr+port,r)
}