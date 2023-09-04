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
		fmt.Printf("Error while reading file: %s", err)
	} else {
		WriteAFile(myWritingPath,data)
		fmt.Print("Complete succesful")
	}
}

// Read a file
// returns data []byte, error
func ReadFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	return data, err
}

// Read a file
// returns data []byte, error
func WriteAFile(path string, data []byte) error {
	// WRITING A FILE
	if err := os.WriteFile(path, data, 0644); err != nil {
		return err
	}
	return nil
}
