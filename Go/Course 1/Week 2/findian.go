package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Write a program which prompts the user to enter a string. The program searches through the entered string for the characters 'i', 'a', and 'n'. The program should print "Found!" if the entered string starts with the character 'i', ends with the character 'n', and contains the character 'a'. The program should print "Not Found!" otherwise. The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.
// Examples: The program should print "Found!" for the following example entered strings, "ian", "Ian", "iuiygaygn", "I d skd a efju N". The program should print "Not Found!" for the following strings, "ihhhhhn", "ina", "xian".
// Submit your source code for the program, "findian.go".

func main() {
	fmt.Println("Enter the string:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	input = strings.ToLower(input)
	input = strings.ReplaceAll(input, " ", "")
	if strings.Contains(input, "a") {
		if strings.HasPrefix(input, "i") {
			if strings.HasSuffix(input, "n") {
				fmt.Printf("Found!")
				os.Exit(0)
			}
		}
	}
	fmt.Print("Not Found!")
}
