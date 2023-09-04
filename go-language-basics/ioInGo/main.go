package main

import (
	"fmt"
	"os"
)

func main() {
	myReadingPath := "/home/sduarte/.config/i3/config"
	myWritingPath := "/home/sduarte/test2"

	data, err := ReadFile(myReadingPath)
	if err != nil {
		fmt.Printf("Error reading the file: %s", err)
		return
	}

	if  err := WriteAFile(myWritingPath,data); err != nil {
		fmt.Printf("Error writing the file: %s", err)
		return
	}

	fmt.Print("Operation completed successfully")
}

// Read a file
func ReadFile(path string) ([]byte, error) {
	// receives a string
	// returns data []byte, error
	data, err := os.ReadFile(path)
	return data, err
}

// Write a file
func WriteAFile(path string, data []byte) error {
	// returns error if not succed
	if err := os.WriteFile(path, data, 0644); err != nil {
		return err
	}
	return nil
}
