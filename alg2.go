package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, World!"
	count := countVowels(str)
	fmt.Println("Quantidade de vogais:", count)
}

func countVowels(str string) int {
	vowels := "aeiouAEIOU"
	count := 0
	for _, char := range str {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}
	return count
}
