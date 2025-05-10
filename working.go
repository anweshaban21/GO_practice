package main

import (
	"fmt"
	"os"
)

func main() {
	// Create a new file
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // Ensure the file is closed after writing

	fmt.Println("File created successfully")

	// Write content to the file
	content := "Hello, World!"
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Content written successfully")
}
