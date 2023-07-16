package main

import "fmt"

func bubblesort(s []int) {
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				swap(s, j)
			}
		}
	}
}

func swap(s []int, i int) {
	var temp int
	temp = s[i]
	s[i] = s[i+1]
	s[i+1] = temp
}

func main() {
	len := 0
	fmt.Print("Enter the number of integers: ")
	fmt.Scanln(&len)
	input := make([]int, len)
	for i := 0; i < len; i++ {
		fmt.Print("Enter a number: ")
		fmt.Scanf("%d", &input[i])
	}
	bubblesort(input)
	fmt.Println("Sorted Slice:", input)
}
