package services

import (
	"todo-list/models"
)

func GetAllTodos() []models.TodoList {
	todoLists := []models.TodoList{
		{TaskName: "Learn Go", Tag: "Programming", Done: false},
		{TaskName: "Walk the dog", Tag: "Personal", Done: true},
		{TaskName: "Read a book", Tag: "Leisure", Done: false},
	}
	return todoLists
}
