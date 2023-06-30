package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, World!"
	newStr, err := replaceVowels(str)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Nova string:", newStr)
	}
}

func replaceVowels(str string) (string, error) {
	if str == "" {
		return "", fmt.Errorf("string vazia")
	}
	vowels := "aeiouAEIOU"
	for _, vowel := range vowels {
		str = strings.ReplaceAll(str, string(vowel), "*")
	}
	return str, nil
}
