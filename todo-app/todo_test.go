package todo_test

import (
	"os"
	"testing"

	"github.com/TaskMasterErnest/todo"
)

// TestAdd tests the Add method of the List type
func TestAdd(t *testing.T) {
	// taking the list of items fro the todo package
	l := todo.List{}

	taskName := "New Task"

	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("expected %q, got %q instead.\n", taskName, l[0].Task)
	}
}

// TestComplete test the Complete method of the list type
func TestComplete(t *testing.T) {
	l := todo.List{}

	taskName := "Completed Task"

	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.\n", taskName, l[0].Task)
	}

	// check if Done = true
	// default is false
	if l[0].Done {
		t.Errorf("Task should not have been completed.")
	}

	// run the complete command
	l.Complete(1)

	// check if task has been completed
	// Done should NOT be false.
	if !l[0].Done {
		t.Errorf("New task should have been completed")
	}
}

// TestDelete tests the Delete method of the list type
func TestDelete(t *testing.T) {
	l := todo.List{}

	tasks := []string{
		"New Task 1",
		"New Task 2",
		"New Task 3",
	}

	// use a loop to add the tasks to the List
	for _, value := range tasks {
		l.Add(value)
	}

	// compare if the tasks out here match the tasks put in the List
	if l[0].Task != tasks[0] {
		t.Errorf("expected %q, got %q instead.\n", tasks[0], l[0].Task)
	}

	// if tasks match, proceed to delete task 2
	l.Delete(2)

	// check whether the tasks in the List have reduced in number
	if len(l) != 2 {
		t.Errorf("expected %d, got %d instead.\n", 2, len(l))
	}

	// the change must be that the second task in List
	// must now match the third task in tasks provided outside
	// adjust for 0-index matching
	if l[1].Task != tasks[2] {
		t.Errorf("expected %q, got %q instead.\n", tasks[2], l[1].Task)
	}
}

// TestSaveAndGet tests the Save and Get methods on the List
func TestSaveAndGet(t *testing.T) {
	// create two lists
	l1 := todo.List{}
	l2 := todo.List{}

	// create a new task in the first list
	taskName := "Task in List 1"

	l1.Add(taskName)

	if l1[0].Task != taskName {
		t.Errorf("expected %q, got %q instead.\n", taskName, l1[0].Task)
	}

	// create a temp file to save the task in
	tf, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("error creating the temp file: %s", err)
	}
	// clean up the file after working with it
	defer os.Remove(tf.Name())

	// call the Save method to work with the new file
	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("error saving list to file: %s", err)
	}

	// call the Get method to retrieve the task in the file
	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("error getting list from file: %s", err)
	}

	// check whether content in both lists match
	if l1[0].Task != l2[0].Task {
		t.Errorf("task %q should match %q task", l1[0].Task, l2[0].Task)
	}
}
