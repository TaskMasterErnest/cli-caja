package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/TaskMasterErnest/todo"
)

// hardcode the name of the file to store the data for the todo list in
const todoFileName = ".todo.json"

func main() {
	// create a pointer to the memory address of an empty instance of the todo.List interface
	l := &todo.List{}

	// call the Get method using this function to check if any data will be gotten back
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// check the number of arguments added to the command and use that to decide
	// what actions to perform, using a switch case
	switch {
	// for no extra arguments, print the list of tasks
	case len(os.Args) == 1:
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	// for no command-line arguments added, concatenate all provided arguments with a space
	// and add them to the list as a Task
	default:
		// concatenate with a space
		item := strings.Join(os.Args[1:], " ")
		// call the Add method
		l.Add(item)
		// then save the task item to the List
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
