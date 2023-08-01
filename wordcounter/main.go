package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// define a boolean flag to count lines instead of words
	lines := flag.Bool("l", false, "count lines")
	// define a boolean flag to count bytes instead of words or lines
	bites := flag.Bool("b", false, "count bytes")
	// parse the command-line into the defined flag
	flag.Parse()

	// call the count function to count number of words
	// count function takes input from Stdin and prints it out
	// UPDATE: adding a pointer to take into account the result of the lines flag passed on the command-line
	fmt.Println(count(os.Stdin, *lines, *bites))
}

// count function takes in a Reader and returns an int
func count(r io.Reader, countLines bool, countBytes bool) int {
	// splitting the words in a line of text
	// call the bufio.ScanWords to do that
	// UPDATE: by default the NewScanner defaults to ScanLines which scans a line of text till the \n delimiter.
	scanner := bufio.NewScanner(r)

	// a conditional that switches between scanning words and scanning lines
	// default is splitting by lines
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	// defining the wordcount
	wordCount := 0

	// scan the words and increment the scanner
	if countBytes {
		for scanner.Scan() {
			wordCount += len(scanner.Bytes())
		}
	} else {
		for scanner.Scan() {
			wordCount++
		}
	}

	err := scanner.Err()
	if err != nil {
		fmt.Printf("error scanning by words %s\n", err)
	}

	return wordCount
}
