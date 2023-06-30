package main

import (
	"fmt"
	"math"
)

func main() {
	number := 17
	isPrime, err := checkPrime(number)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("É primo:", isPrime)
	}
}

func checkPrime(number int) (bool, error) {
	if number < 0 {
		return false, fmt.Errorf("número negativo")
	}
	if number < 2 {
		return false, nil
	}
	sqrt := int(math.Sqrt(float64(number)))
	for i := 2; i <= sqrt; i++ {
		if number%i == 0 {
			return false, nil
		}
	}
	return true, nil
}
