package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// call the count function to count number of words
	// count function takes input from Stdin and prints it out
	fmt.Println(count(os.Stdin))
}

// count function takes in a Reader and returns an int
func count(r io.Reader) int {
	// splitting the words in a line of text
	// call the bufio.ScanWords to do that
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	// defining the wordcount
	wordCount := 0

	// scan the words and increment the scanner
	for scanner.Scan() {
		wordCount++
	}

	err := scanner.Err()
	if err != nil {
		fmt.Printf("error scanning by words %s\n", err)
	}

	return wordCount
}
