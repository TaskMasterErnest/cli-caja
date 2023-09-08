package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const (
	defaultTemplate = `
<!DOCTYPE html>
<html>
  <head>
    <meta http-equiv="content-type" content="text/html; charset=utf-8">
    <title>{{ .Title }}</title>
  </head>
  <body>
    {{ .Body }}
  </body>
</html>
`
)

// a content type that defines the HTML content to add into the template
type content struct {
	Title string
	Body  template.HTML
}

func main() {
	// Parse in flags to take in file that contains the markdown to be used
	filename := flag.String("file", "", "Markdown file to preview")
	skipPreview := flag.Bool("s", false, "Skip auto-preview")
	tFName := flag.String("t", "", "Alternate template file")
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
	if err := run(*filename, *tFName, os.Stdout, *skipPreview); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// The run() function takes in a single value, the filename
// it reads the data from the file, parses the content to HTML and returns the content wrapped in HTML
// include an io Interface to capture the output to be used for the integration test
func run(filename string, tFName string, out io.Writer, skipPreview bool) error {
	// read the file from the filename given
	input, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// the parsecontent() function converts the Markdown to HTML
	htmlData, err := parsecontent(input, tFName)
	if err != nil {
		return err
	}

	// create a temporary file to write the processed data to
	temp, err := os.CreateTemp("", "mdp*.html")
	if err != nil {
		return err
	}
	// close the file after writing to it
	if err := temp.Close(); err != nil {
		return err
	}

	// assign name of temp file outName
	outName := temp.Name()

	fmt.Fprintln(out, outName)

	// the saveHTML() function runs to save the final converted content to a filename
	if err := saveHTML(outName, htmlData); err != nil {
		return err
	}

	if skipPreview {
		return nil
	}

	// remove the files
	os.Remove(outName)

	return preview(outName)
}

// the parsecontent() function takes after the blackfriday package, it uses the run function to perform the conversion
// it takes in the data as bytes and returns bytes
func parsecontent(input []byte, tFName string) ([]byte, error) {
	// using the blackfriday Run() function to perform the conversion
	output := blackfriday.Run(input)
	// using bluemonday to sanitize the output
	body := string(bluemonday.UGCPolicy().SanitizeBytes(output))

	// Parse content of defaultTemplate const into the Template
	t, err := template.New("mdp").Parse(defaultTemplate)
	if err != nil {
		return nil, err
	}

	// check if the user has defined any custom templates to be used
	if tFName != "" {
		t, err = template.ParseFiles(tFName)
		if err != nil {
			return nil, err
		}
	}

	// instantiate the content tyep and force the Body to be converted using template.HTML
	c := content{
		Title: "Markdown Preview Tool",
		Body:  template.HTML(body),
	}

	// Now join the contents of the newly created HTML body using a buffer of bytes
	// create a storage buffer to write to file
	var buffer bytes.Buffer

	// write data to the buffer by executing the template, with the content type
	if err := t.Execute(&buffer, c); err != nil {
		return nil, err
	}

	// return the buffer
	return buffer.Bytes(), nil
}

// the saveHTML() function, which will receive th eentire HTML content stored in the buffer
// and save it to a file specifiedby the parameter name outFName
func saveHTML(outFName string, data []byte) error {
	// write the bytes to the file
	return os.WriteFile(outFName, data, 0644)
}

// add a preview function to automatically open files when converted
func preview(fname string) error {
	cName := ""
	cParams := []string{}

	// define the executable file based on the OS
	switch runtime.GOOS {
	case "Linux":
		cName = "xdg-open"
	case "Windows":
		cName = "cmd.exe"
		cParams = []string{"/C", "start"}
	case "darwin":
		cName = "open"
	default:
		return fmt.Errorf("OS not supported")
	}

	// append the filename to the params slice
	cParams = append(cParams, cName)

	// locate the executable file in the path
	cPath, err := exec.LookPath(cName)
	if err != nil {
		return err
	}

	// open the file using the default program
	err = exec.Command(cPath, cParams...).Run()

	// allow the browser time to open the file before deleting it
	time.Sleep(5 * time.Second)
	return err
}
