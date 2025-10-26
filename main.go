package main

import (
	database "github.com/Tola-lemma/golang_todo/Database"
	migrations "github.com/Tola-lemma/golang_todo/Migrations"
)

func main() {
	database.InitDB()
	migrations.RunMigrations()
}