package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	Database "github.com/Tola-lemma/golang_todo/Database"
	models "github.com/Tola-lemma/golang_todo/Models"
	"github.com/Tola-lemma/golang_todo/Validations"
	"github.com/gorilla/mux"
)

func GetTodos(w http.ResponseWriter, r *http.Request){
	query:="SELECT id, title, completed, created_at FROM todos ORDER BY id DESC"
	rows, err :=Database.DB.Query(query)
	if err !=nil{
		http.Error(w , err.Error(),http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var todos []models.Todo
	for rows.Next(){
		var t models.Todo
		if err:=rows.Scan(&t.ID,&t.Title,&t.Completed,&t.CreatedAt); err!=nil{
			http.Error(w,err.Error(),http.StatusNotFound)
			return
		}
		todos=append(todos, t)
	}
	json.NewEncoder(w).Encode(todos)
}
// --- CREATE TODO ---
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var input Validations.TodoInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	if err := Validations.ValidateTodoInput(input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := Database.DB.Exec("INSERT INTO todos (title) VALUES ($1)", input.Title)
	if err != nil {
		http.Error(w, "Failed to insert todo", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Todo created successfully"})
}

// --- UPDATE TODO ---
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var input Validations.TodoUpdateInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	if err := Validations.ValidateTodoUpdateInput(input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
 //dynamic update
 query :="UPDATE todos SET "
 args :=[]any{}
 argPos:=1
 if input.Title !=nil{
	query +=fmt.Sprintf("title=$%d,",argPos)
	args=append(args, *input.Title)
	argPos++
 }
 if input.Completed !=nil{
	query +=fmt.Sprintf("completed=$%d,",argPos)
	args=append(args, *input.Completed)
	argPos++
 }
	//remove trailing comma
	if len(args) ==0{
		http.Error(w,"No fields to Update",http.StatusBadRequest)
		return
	}
	query = query[:len(query)-1] //remove trailing comma
	query +=fmt.Sprintf(" WHERE id=$%d",argPos)
	args=append(args, id)
    result, err := Database.DB.Exec(query, args...)
	if err != nil {
		http.Error(w, "Failed to update todo", http.StatusInternalServerError)
		return
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Todo updated successfully"})
}

// --- DELETE TODO ---
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	result, err := Database.DB.Exec("DELETE FROM todos WHERE id=$1", id)
	if err != nil {
		http.Error(w, "Failed to delete todo", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Todo deleted successfully"})
}
