package main

import (
	"fmt"
	"math"
)

func main() {
	number := 10
	primes, err := findPrimes(number)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Números primos:", primes)
	}
}

func findPrimes(number int) ([]int, error) {
	if number < 0 {
		return nil, fmt.Errorf("número negativo")
	}
	primes := []int{}
	for i := 2; i <= number; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes, nil
}

func isPrime(number int) bool {
	if number < 2 {
		return false
	}
	sqrt := int(math.Sqrt(float64(number)))
	for i := 2; i <= sqrt; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
