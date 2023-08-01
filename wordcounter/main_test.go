package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	// create a buffer, a space to hold stuff, that takes in a words but stores them as bytes
	bytesInBuffer := bytes.NewBufferString("word1 word2 word3 word4 word5\n")

	// register the number of bytes in buffer
	numberOfWords := 5

	// call the count function to count the words in the buffer
	// UPDATE: add false as input as response to the flag to be passed
	result := count(bytesInBuffer, false, false)

	// compare the results
	if result != numberOfWords {
		t.Errorf("expected %d, got %d instead.\n", numberOfWords, result)
	}
}

func TestCountLines(t *testing.T) {
	// create a buffer to take in lines of text and store them as bytes
	bytesInBuffer := bytes.NewBufferString("word1 word2 word3\nline1\nline2 word4")

	// register the number of lines
	numberOfLines := 3

	// call the count function with the line flag passed to count number of lines in the buffer
	result := count(bytesInBuffer, true, false)

	// compare the results
	if result != numberOfLines {
		t.Errorf("expected %d, got %d instead.\n", numberOfLines, result)
	}
}

func TestCountBytes(t *testing.T) {
	// create a buffer to take in words and hold them as bytes
	bytesInBuffer := bytes.NewBufferString("word1 word2 word3 word4 word5 word6")

	// register the number of bytes
	numberOfBytes := 35

	// call the count function to count the number of bytes with the bytes flag
	result := count(bytesInBuffer, false, true)

	// compare the results
	if result != numberOfBytes {
		t.Errorf("expected %d, got %d instead.\n", numberOfBytes, result)
	}
}
