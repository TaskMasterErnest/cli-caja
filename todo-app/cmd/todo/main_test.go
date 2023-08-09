package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// names of the binary and filenames that will be used in this test
var (
	binName  = "todo"
	fileName = ".todo.json"
)

// For this test, we will do the following
// 1. compile the application using the go build tool into a binary
// 2. execute the binary with different arguments and ascertain the correct behaviour of the application
// the important packages to use are the TestMain and the os/exec packages

// create TestMain to call the go build tool, build the binary, execute the the tests and clean up after it is done
func TestMain(m *testing.M) {
	fmt.Println("Building tool ...")

	// check the OS runtime, default is Linux
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	// build the application
	build := exec.Command("go", "build", "-o", binName)

	// Run the build and check for errors
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		os.Exit(1)
	}

	fmt.Println("Running tests ...")
	result := m.Run()

	// cleaning up the files
	fmt.Println("Cleaning up ...")
	os.Remove(binName)
	os.Remove(fileName)

	os.Exit(result)
}

// Testing the tool against the binary created using subtests
// These tests will depend on each other by utilizing the t.Run method of the testing package
func TestTodoCLI(t *testing.T) {
	// specify a task to input
	task := "Adding a new task 1"

	// specify the dir, since the TestMain compiles the binary in the same dir
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	// specify the command path to use to execute the binary
	cmdPath := filepath.Join(dir, binName)

	// create the first test, which ensures that the tool can add a new task
	// use the subtest t.Run function
	t.Run("AddNewtask", func(t *testing.T) {
		cmd := exec.Command(cmdPath, strings.Split(task, " ")...) // split the task by the spaces and pass then in one by one

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	// create a second test, which ensure that the tool can list the tasks
	t.Run("ListTasks", func(t *testing.T) {
		// by default, running the command without arguments shold list the tasks
		cmd := exec.Command(cmdPath)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		// specify the result expected
		expected := task + "\n"

		if expected != string(output) {
			t.Errorf("expected %q, got %q instead \n", expected, string(output))
		}
	})
}
