package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	target := 3
	index := findElementIndex(numbers, target)
	fmt.Println("Índice do elemento:", index)
}

func findElementIndex(numbers []int, target int) int {
	for i, num := range numbers {
		if num == target {
			return i
		}
	}
	return -1
}
