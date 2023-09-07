package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

// remove the hardcoded resultFile variable
const (
	inputFile  = "./testdata/test1.md"
	goldenFile = "./testdata/test1.md.html"
)

func TestParseContent(t *testing.T) {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatal(err)
	}

	result := parsecontent(input)

	expected, err := os.ReadFile(goldenFile)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(expected, result) {
		t.Logf("golden:\n%s\n", expected)
		t.Logf("result:\n%s\n", result)
		t.Error("result content does not match golden file")
	}
}

// integrated test case that tests the run function
func TestRun(t *testing.T) {
	// using a buffer to capture the new name of the generated file
	var mockStdOut bytes.Buffer

	if err := run(inputFile, &mockStdOut); err != nil {
		t.Fatal(err)
	}

	// the resulting resultFile is extracted
	resultFile := strings.TrimSpace(mockStdOut.String())

	//check the result
	result, err := os.ReadFile(resultFile)
	if err != nil {
		t.Fatal(err)
	}

	//check the goldenfile
	expected, err := os.ReadFile(goldenFile)
	if err != nil {
		t.Fatal(err)
	}

	// compare the results
	if !bytes.Equal(expected, result) {
		t.Logf("golden:\n%s\n", expected)
		t.Logf("result:\n%s\n", result)
		t.Error("result content does not match golden file")
	}

	// remove the file after checking
	os.Remove(resultFile)
}
