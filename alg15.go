package main

import (
	"fmt"
)

func main() {
	number := 3
	slice := []int{1, 2, 3, 4, 5}
	isPresent, err := checkNumberPresence(number, slice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Está presente:", isPresent)
	}
}

func checkNumberPresence(number int, slice []int) (bool, error) {
	if len(slice) == 0 {
		return false, fmt.Errorf("slice vazio")
	}
	for _, num := range slice {
		if num == number {
			return true, nil
		}
	}
	return false, nil
}
