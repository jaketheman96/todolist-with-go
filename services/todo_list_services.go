package services

import (
	"todo-list/models"
)

func GetAllTodos() []models.TodoList {
	todoLists := []models.TodoList{
		{TaskName: "Learn Golan", Tag: "", Done: false},
		{TaskName: "Walk the dog", Tag: "", Done: true},
		{TaskName: "Read a book", Tag: "Leisur", Done: false},
	}
	return todoLists
}
