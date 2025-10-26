package database

import (
	"database/sql"
	"fmt"
	"log"
    _ "github.com/lib/pq"   
	"github.com/spf13/viper"
)
var DB *sql.DB
func InitConfig() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err !=nil{
		log.Printf("Config File not found: %v",err)
	}
}
func InitDB(){
	InitConfig()
	connStr:=fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_NAME"),
		viper.GetString("SSL_MODE"),
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