package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/TaskMasterErnest/todo"
)

// making the ToDo filename a variable
var todoFileName = ".todo.json"

// make a function that will print a list, it takes in a flag as a receiver
func expandList(l *todo.List, verbose bool) {
	// adding a formatted list
	formatted := ""
	// loop over the list and print out the tasks
	for index, t := range *l {
		// set a prefix
		prefix := "  "
		if t.Done {
			prefix = "X "
		}
		if verbose {
			if t.Done {
				formatted += fmt.Sprintf("%s%d: %s -- Created: %v, Completed: %v\n", prefix, index+1, t.Task, t.CreatedAt, t.CompletedAt)
			} else {
				formatted += fmt.Sprintf("%s%d: %s -- Created: %v\n", prefix, index+1, t.Task, t.CreatedAt)
			}
		} else {
			formatted += fmt.Sprintf("%s%d: %s\n", prefix, index+1, t.Task)
		}
	}
	fmt.Println(formatted)
}

func main() {
	// implement the use of a ENV_VAR to set the filename to save to
	// check if the user has defined an ENV_VAR for the custom file name
	// name of ENV_VAR should be TODO_FILENAME
	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}

	// adding a usage flag that points to all the functions
	// we add usage information and display a custom message when the code is run
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s tool. Developed for use by Ernest Klu\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright 2023\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage information\n")
		flag.PrintDefaults()
	}

	// add flags to parse and pass into the command-line
	add := flag.Bool("add", false, "Add task to ToDo list") // change this flag to add with Bool
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")
	delete := flag.Int("delete", 0, "Item to be deleted")
	verbose := flag.Bool("verbose", false, "Date/Time to tasks listed")
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
		// list the current lists,
		fmt.Println(l)
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
	case *add: // if the task flag is called and the arguments are not empty,
		// addthe task to the List
		// take any arguments, excluding flags, and use them as the new task
		t, err := GetTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		// add the task to the list
		l.Add(t)
		// then save the data to the List
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	// adding the case to delete a task from the List
	case *delete > 0:
		if err := l.Delete(*delete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		// after deletion save the data into a new compiled lists
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	// adding a case to enable verbose output
	case *verbose:
		expandList(l, *verbose)

	default:
		// we assume an invalid falg was passed in, so we throw an error
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}

// Define a new function that can receive tasks from other sources and not just arguments only
// The GetTask function will take in arguments passed from other sources through standard input
func GetTask(r io.Reader, args ...string) (string, error) {
	// check if any arguments were provided as parameters
	// if there are, concatenate them to form the string that will be used as a Task input
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	// read the input from the standard input as it has been passed
	s := bufio.NewScanner(r)
	// scan the input
	s.Scan()
	// check for errors when reading the input
	if err := s.Err(); err != nil {
		return "", err
	}

	// if no errors, check that the scan is populated with Text than is useful
	if len(s.Text()) == 0 {
		return "", fmt.Errorf("Task should not be blank")
	}

	// if all is well, return the scanned text
	return s.Text(), nil
}
