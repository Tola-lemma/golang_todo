package Validations

import (
	"errors"
	"strings"
)

type TodoInput struct {
	Title string `json:"title"`
}

func ValidateTodoInput(t TodoInput) error {
	t.Title = strings.TrimSpace(t.Title)
	if len(t.Title) < 3 {
		return errors.New("title must be at least 3 characters long")
	}
	return nil
}
