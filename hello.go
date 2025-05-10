package main

import (
	"fmt"
	"os"
)

func main() {
	// 	fmt.Println("Hello, World!")
	// }
	dir, _ := os.Getwd()
	fmt.Println("Current working directory:", dir)
}
