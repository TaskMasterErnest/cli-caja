package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

// the item is local to this package, not to be exported
// item struct to represent the ToDo items
type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

// the list item is global, accessible to users outside of the todo package
// List represent the list of ToDO items
type List []item

// creating methods to interact with the list and item types

// Add creates a new Todo and appends it to the List
func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	// de-reference the pointer to make the underlying value accessible to be used
	*l = append(*l, t)
}

// Complete method marks a ToDo item as completed
// by setting Done = true and CompletedAt = time.Now()
func (l *List) Complete(i int) error {
	listItem := *l
	if i <= 0 || i > len(listItem) {
		return fmt.Errorf("item %d does not exist", i)
	}

	// Adjusting index for 0-based indexing
	listItem[i-1].Done = true
	listItem[i-1].CompletedAt = time.Now()

	return nil
}

// Delete method removes an item from the list
func (l *List) Delete(i int) error {
	listItem := *l
	if i <= 0 || i > len(listItem) {
		return fmt.Errorf("Item %d does not exist", i)
	}

	// restructuring position of items to mark deletion
	*l = append(listItem[:i-1], listItem[i:]...)

	return nil
}

// Adding the Save method which converts the data to JSON and stores it in a filename provided
func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}

	// write to the filename
	return os.WriteFile(filename, js, 0644)
}

// Get method takes in a filename, opens it, reads it, decodes the JSON data and parses it into a list
func (l *List) Get(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		// check if the file exists
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}

	return json.Unmarshal(file, l)
}
