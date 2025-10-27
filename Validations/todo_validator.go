package Validations

import (
	"errors"
	"strings"
)

type TodoInput struct {
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

func ValidateTodoInput(t TodoInput) error {
	t.Title = strings.TrimSpace(t.Title)
	if len(t.Title) < 3 {
		return errors.New("title must be at least 3 characters long")
	}
	return nil
}
//update optional fields
type TodoUpdateInput struct{
	Title *string `json:"title,omitempty"`
	Completed *bool `json:"completed,omitempty"`
}
func ValidateTodoUpdateInput(t TodoUpdateInput) error{
	if t.Title!=nil && len(*t.Title)<3{
       return errors.New("title must be at least 3 characters long")
	}
	return  nil
}