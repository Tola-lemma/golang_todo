package config

import (
	"fmt"
	"log"
	"github.com/spf13/viper"
)

type Config struct {
	DBUser        string `mapstructure:"DB_USER"`
	DBPass        string `mapstructure:"DB_PASSWORD"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBName        string `mapstructure:"DB_NAME"`
	SSLMode       string `mapstructure:"SSL_MODE"`
	ServerPort    string `mapstructure:"SERVER_PORT"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

var AppConfig Config

func LoadConfig() {
	viper.AddConfigPath(".")       
	viper.SetConfigFile(".env")    
	viper.AutomaticEnv()     

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Unable to decode config into struct: %v", err)
	}

	fmt.Println("Config loaded successfully")
}
