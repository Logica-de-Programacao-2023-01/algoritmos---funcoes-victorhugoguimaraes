package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	evens, err := filterEvenNumbers(numbers)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Números pares:", evens)
	}
}

func filterEvenNumbers(numbers []int) ([]int, error) {
	if len(numbers) == 0 {
		return nil, fmt.Errorf("slice vazio")
	}
	evens := []int{}
	for _, num := range numbers {
		if num%2 == 0 {
			evens = append(evens, num)
		}
	}
	return evens, nil
}
