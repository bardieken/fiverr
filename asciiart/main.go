package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Check that the user has provided input
	if len(os.Args) <= 1 {
		fmt.Println("Please enter some text")
		return
	}
	input1 := os.Args
	if input1[1] == "" {
		os.Exit(0)
	}
	// Read the contents of the file "standard.txt"
	file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Printf("Error message: %s:\n", err)
		os.Exit(2)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error!")
		os.Exit(2)
	}
	lines := strings.Split(string(data), "\n")
	// Convert the user input to a slice of runes
	input := []rune(os.Args[1])
	// If the input is "\n", print a newline and exit
	if input[0] == '\\' && input[1] == 'n' {
		fmt.Println("")
		os.Exit(0)
	}
	// Nested loop to print line by line depending on input
	words := strings.Split(string(input), "\\n")
	for _, word := range words {
		for h := 1; h < 9; h++ {
			if word == "" {
				fmt.Println("")
				break
			}
			for _, l := range []byte(word) {
				for i, line := range lines {
					if i == (int(l)-32)*9+h {
						fmt.Print(line)
					}
				}
			}
			fmt.Println()
		}
	}
}
