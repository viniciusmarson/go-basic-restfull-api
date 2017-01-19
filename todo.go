package main

import "time"

//Todo will store the todo items
type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

//Todos will store todos
type Todos []Todo
