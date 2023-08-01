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
	result := count(bytesInBuffer)

	// compare the results
	if result != numberOfBytes {
		t.Errorf("expected %d, got %d instead \n", numberOfBytes, result)
	}
}
