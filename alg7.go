package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	doubledSlice, err := applyFunction(numbers, double)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Novo slice:", doubledSlice)
	}
}

func double(num int) int {
	return num * 2
}

func applyFunction(numbers []int, fn func(int) int) ([]int, error) {
	if len(numbers) == 0 {
		return nil, fmt.Errorf("slice vazio")
	}
	result := make([]int, len(numbers))
	for i, num := range numbers {
		result[i] = fn(num)
	}
	return result, nil
}
