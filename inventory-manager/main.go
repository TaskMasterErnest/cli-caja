package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/cli-caja/inventory-manager/inventory"
)

func CheckFile(filename string) error {
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("File: %s does not exist\n", filename)
			os.Exit(1)
		}
	}
	return nil
}

func GetArgs(r io.Reader, args ...string) ([]string, error) {
	if len(args) > 0 {
		return strings.Split(strings.Join(args, " "), " "), nil
	}
	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return nil, err
	}

	if len(s.Text()) == 0 {
		return nil, fmt.Errorf("content should not be blank")
	}

	return strings.Split(s.Text(), " "), nil
}

func main() {
	// files
	inventoryFile := flag.String("f", "", "Path to the inventory file")
	// actions
	list := flag.Bool("l", false, "List the contents of file")
	add := flag.Bool("add", false, "Add record to file")
	//parser
	flag.Parse()

	// require the files to be present
	if *inventoryFile == "" {
		fmt.Println("Inventory file location is required")
		flag.PrintDefaults()
		os.Exit(1)
	}

	CheckFile(*inventoryFile)
	AddNewLineIfMissing(*inventoryFile)

	// work on inventory file
	file, err := os.Open(*inventoryFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer file.Close()

	// start actions
	switch {
	case *list:
		list := inventory.ListContent(file)
		fmt.Println()
		for _, l := range list {
			fmt.Printf("%v\n", l)
		}
	case *add:
		add, err := GetArgs(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		err = inventory.AddContent(*inventoryFile, add)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error adding content:", err)
			os.Exit(1)
		}
		fmt.Println("Content added successfully!")
	}

}

// prepare the CSV file for appending data
func AddNewLineIfMissing(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	if fileInfo.Size() == 0 {
		return nil
	}

	// read the last character
	reader := bufio.NewReader(file)
	_, err = file.Seek(-1, io.SeekEnd)
	if err != nil {
		return err
	}
	lastChar, err := reader.ReadByte()
	if err != nil {
		return err
	}

	if lastChar != '\n' {
		// Open the file in append mode to add a newline
		file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = file.WriteString("\n")

		if err != nil {
			return err
		}
	}

	return nil

}
