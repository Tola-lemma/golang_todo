package migrations

import (
	"log"

	database "github.com/Tola-lemma/golang_todo/Database"
	schema "github.com/Tola-lemma/golang_todo/Schema"
)

func RunMigrations() {
	tables := []string{
		schema.CreateTodoTable,
		//others go here
	}
	for _, query:= range tables{
		_, err := database.DB.Exec(query)
		if err !=nil{
			log.Fatalf("Migration failed : %v",err)
		}
	}
	log.Println("Migrations Executed Successfully!")
}