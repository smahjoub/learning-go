package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("File path is missing \n")
		os.Exit(1)
	}
	filePath := os.Args[1]

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error:  %s", err)
		os.Exit(1)
	}

	fmt.Printf("File contents: %s", content)

}
