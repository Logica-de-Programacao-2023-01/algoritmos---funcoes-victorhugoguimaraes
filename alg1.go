package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	average := calculateAverage(numbers)
	fmt.Println("Média:", average)
}

func calculateAverage(numbers []int) float64 {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	average := float64(sum) / float64(len(numbers))
	return average
}
