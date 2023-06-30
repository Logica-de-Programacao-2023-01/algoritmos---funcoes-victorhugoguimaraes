package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{5, 2, 4, 1, 3}
	sortedSlice, err := sortSlice(numbers)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Slice ordenado:", sortedSlice)
	}
}

func sortSlice(numbers []int) ([]int, error) {
	if len(numbers) == 0 {
		return nil, fmt.Errorf("slice vazio")
	}
	sorted := make([]int, len(numbers))
	copy(sorted, numbers)
	sort.Ints(sorted)
	return sorted, nil
}
