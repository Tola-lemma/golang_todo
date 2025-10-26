package controller

import (
	"encoding/json"
	"net/http"

	database "github.com/Tola-lemma/golang_todo/Database"
	models "github.com/Tola-lemma/golang_todo/Models"
)

func GetTodos(w http.ResponseWriter, r *http.Request){
	query:="SELECT id, title, completed, created_at FROM todos ORDER BY id DESC"
	rows, err :=database.DB.Query(query)
	if err !=nil{
		http.Error(w , err.Error(),http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var todos []models.Todo
	for rows.Next(){
		var t models.Todo
		if err:=rows.Scan(&t.ID,&t.Title,&t.Completed,&t.Completed); err!=nil{
			http.Error(w,err.Error(),http.StatusNotFound)
			return
		}
		todos=append(todos, t)
	}
	json.NewEncoder(w).Encode(todos)
}