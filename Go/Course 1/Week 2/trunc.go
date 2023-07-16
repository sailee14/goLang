package main

import (
	"fmt"
	"log"
)

// Write a program which prompts the user to enter a floating point number and prints the integer which is a truncated version of the floating point number that was entered. Truncation is the process of removing the digits to the right of the decimal place.
// Submit your source code for the program, "trunc.go".

func main() {
	var inputNumber float64
	fmt.Println("Enter the float value:")
	_, err := fmt.Scanln(&inputNumber)
	if err != nil {
		log.Printf("[Error] Invalid Input!")
	}
	fmt.Printf("Truncated version is %v", int64(inputNumber))
}
