package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{5, 2, 9, 1, 7, 3}
	secondLargest := findSecondLargest(numbers)
	fmt.Println("Segundo maior valor:", secondLargest)
}

func findSecondLargest(numbers []int) int {
	sort.Ints(numbers)
	if len(numbers) >= 2 {
		return numbers[len(numbers)-2]
	}
	return -1
}
