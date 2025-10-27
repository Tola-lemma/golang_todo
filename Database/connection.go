package database

import (
	"database/sql"
	"fmt"
	"log"
    _ "github.com/lib/pq"   
config "github.com/Tola-lemma/golang_todo/Config"
)
var DB *sql.DB
func InitDB(){
	cfg := config.AppConfig
	connStr:=fmt.Sprintf(
	"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.SSLMode,
	)
	var err error
	DB, err = sql.Open("postgres",connStr)
	 if err !=nil {
		log.Fatalf("Failed to Connect to DB: %v",err)
	 }
	 if err = DB.Ping(); err !=nil{
		log.Fatalf("DB not reachable: %v",err)
	 }
	 log.Println("Connected to PostgreSQL!")
}