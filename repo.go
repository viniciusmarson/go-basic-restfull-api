package main

import "fmt"

var currentID int

var todos Todos

// Give us some seed data
func init() {
	RepoCreateTodo(Todo{Name: "Write presentation"})
	RepoCreateTodo(Todo{Name: "Host meetup"})
}

//RepoFindTodo will find a Todo item by the id
func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.ID == id {
			return t
		}
	}
	// return empty Todo if not found
	return Todo{}
}

//RepoCreateTodo will create a Todo item
func RepoCreateTodo(t Todo) Todo {
	currentID++
	t.ID = currentID
	todos = append(todos, t)
	return t
}

//RepoDestroyTodo will destroy all Todo items in the list
func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
