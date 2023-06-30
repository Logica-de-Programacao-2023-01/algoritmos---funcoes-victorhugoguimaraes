package main

import (
	"fmt"
)

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{4, 5, 6, 7, 8}
	commonNumbers, err := findCommonNumbers(slice1, slice2)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Números comuns:", commonNumbers)
	}
}

func findCommonNumbers(slice1, slice2 []int) ([]int, error) {
	if len(slice1) == 0 || len(slice2) == 0 {
		return nil, fmt.Errorf("um dos slices está vazio")
	}
	uniqueNumbers := make(map[int]bool)
	commonNumbers := []int{}
	for _, num := range slice1 {
		uniqueNumbers[num] = true
	}
	for _, num := range slice2 {
		if uniqueNumbers[num] {
			commonNumbers = append(commonNumbers, num)
		}
	}
	return commonNumbers, nil
}
