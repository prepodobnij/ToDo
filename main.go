package main

import (
	"ToDo/todo"
	"fmt"
)

func main() {
	commandsMap := map[string]string{
		"add":      "Add a new task to the todo list",
		"complete": "Set a task as being complete",
		"del":      "Removes a task for the todo list by its id",
		"list":     "List all of the tasks in your todo list",
		"help":     "Help about any commands"}

	commands := [5]string{"add", "complete", "del", "list", "help"}

	fmt.Println("Welcome to the ToDo!!!")
	for {
		todo.HandleCommand(commandsMap, commands)

	}

}
