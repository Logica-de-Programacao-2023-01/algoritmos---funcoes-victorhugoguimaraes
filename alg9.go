package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, World!"
	words, err := extractWords(str)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Palavras:", words)
	}
}

func extractWords(str string) ([]string, error) {
	if str == "" {
		return nil, fmt.Errorf("string vazia")
	}
	words := strings.Fields(str)
	return words, nil
}
