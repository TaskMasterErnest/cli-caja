package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	// create a buffer, a space to hold stuff, that takes in a words but stores them as bytes
	bytesInBuffer := bytes.NewBufferString("word1 word2 word3 word4 word5\n")

	// register the number of bytes in buffer
	numberOfBytes := 5

	// call the count function to count the words in the buffer
	// UPDATE: add false as input as response to the flag to be passed
	result := count(bytesInBuffer, false)

	// compare the results
	if result != numberOfBytes {
		t.Errorf("expected %d, got %d instead.\n", numberOfBytes, result)
	}
}

func TestCountLines(t *testing.T) {
	// create a buffer to take in lines of text and store them as bytes
	bytesInBuffer := bytes.NewBufferString("word1 word2 word3\nline1\nline2 word4")

	// register the number of lines
	numberOfLines := 3

	// call the count function with the line flag passed to count number of lines in the buffer
	result := count(bytesInBuffer, true)

	// compare the results
	if result != numberOfLines {
		t.Errorf("expected %d, got %d instead.\n", numberOfLines, result)
	}
}
