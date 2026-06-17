package main

import (
	"fmt"
	"errors"
)

func main() {
	funcA()
}

func funcA() (string, error) {
	var err error
	filename := ""
	data := ""

	filename, err = GetFileName()
	if err != nil {
		fmt.Println("Error getting filename:", err)
		goto Done
	}

	data, err = ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		goto Done
	}
	fmt.Println("File data:", data)

Done:
	return data, err
}

func GetFileName() (string, error) {
	return "sample.txt", nil
}

func ReadFile(filename string) (string, error) {
	return "Hello, World!", errors.New("file not found")
}