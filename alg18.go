package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	sum, err := applyAndSum(numbers, square)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Soma:", sum)
	}
}

func square(num int) int {
	return num * num
}

func applyAndSum(numbers []int, fn func(int) int) (int, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("slice vazio")
	}
	sum := 0
	for _, num := range numbers {
		sum += fn(num)
	}
	return sum, nil
}
