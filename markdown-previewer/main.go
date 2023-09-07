package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const (
	header = `
<!DOCTYPE html>
<html>
  <head>
    <meta http-equiv="content-type" content="text/html; charset=utf-8">
    <title>Markdown Preview Tool</title>
  </head>
  <body>
`

	footer = `
  </body>
</html>
  `
)

func main() {
	// Parse in flags to take in file that contains the markdown to be used
	filename := flag.String("file", "", "Markdown file to preview")
	// parse the content
	flag.Parse()

	// check if the file has been specified by user
	// if not, call the Usage function and exit
	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}

	// if the filename is present, call a function run() to take in the data
	// if data errors out, send it to Stderr
	if err := run(*filename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// The run() function takes in a single value, the filename
// it reads the data from the file, parses the content to HTML and returns the content wrapped in HTML
func run(filename string) error {
	// read the file from the filename given
	input, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// the parsecontent() function converts the Markdown to HTML
	htmlData := parsecontent(input)

	// set the name and path for the output file for final converted content
	// print out the name of the final file name
	outName := fmt.Sprintf("%s.html", filepath.Base(filename))
	fmt.Println(outName)

	// the saveHTML() function runs to save the final converted content to a filename
	return saveHTML(outName, htmlData)
}

// the parsecontent() function takes after the blackfriday package, it uses the run function to perform the conversion
// it takes in the data as bytes and returns bytes
func parsecontent(input []byte) []byte {
	// using the blackfriday Run() function to perform the conversion
	output := blackfriday.Run(input)
	// using bluemonday to sanitize the output
	body := string(bluemonday.UGCPolicy().SanitizeBytes(output))

	// Now join the contents of the newly created HTML body using a buffer of bytes
	// create a storage buffer to write to file
	var buffer bytes.Buffer

	// write HTML to bytes buffer
	buffer.WriteString(header)
	buffer.WriteString(body)
	buffer.WriteString(footer)

	// return the buffer
	return buffer.Bytes()
}

// the saveHTML() function, which will receive th eentire HTML content stored in the buffer
// and save it to a file specifiedby the parameter name outFName
func saveHTML(outFName string, data []byte) error {
	// write the bytes to the file
	return os.WriteFile(outFName, data, 0644)
}
