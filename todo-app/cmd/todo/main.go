package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/TaskMasterErnest/todo"
)

// hardcode the name of the file to store the data for the todo list in
const todoFileName = ".todo.json"

func main() {
	// add flags to parse ad pass into the command-line
	task := flag.String("task", "", "Task to be included in the ToDo string")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")
	// after stating the flags, parse them in so that they can be used
	// note that in order to use them in this state, they are pointers hence have to be dereferenced by a *
	flag.Parse()

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
	case *list:
		for _, item := range *l {
			if !item.Done { // exclude the completed items from being listed
				fmt.Println(item.Task)
			}
		}
	case *complete > 0: // check if the number given is completed
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		// if the item is found and declared as complete, save the new data into the file
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *task != "": // if the task flag is called and the arguments are not empty,
		// addthe task to the List
		l.Add(*task)
		// then save the data to the List
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		// we assume an invalid falg was passed in, so we throw an error
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}
