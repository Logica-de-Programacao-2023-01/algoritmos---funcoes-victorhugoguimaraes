package main

import (
	"fmt"
	"strconv"
)

func main() {
	number := 12345
	sum, err := sumDigits(number)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Soma dos dígitos:", sum)
	}
}

func sumDigits(number int) (int, error) {
	if number < 0 {
		return 0, fmt.Errorf("número negativo")
	}
	str := strconv.Itoa(number)
	sum := 0
	for _, char := range str {
		digit, _ := strconv.Atoi(string(char))
		sum += digit
	}
	return sum, nil
}
