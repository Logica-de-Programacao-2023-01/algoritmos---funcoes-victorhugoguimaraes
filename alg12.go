package main

import (
	"fmt"
	"unicode"
)

func main() {
	strings := []string{"Hello", "world", "Go"}
	result, err := filterUpperCase(strings)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Strings filtradas:", result)
	}
}

func filterUpperCase(strings []string) (string, error) {
	if len(strings) == 0 {
		return "", fmt.Errorf("slice vazio")
	}
	result := ""
	for _, str := range strings {
		firstChar := []rune(str)[0]
		if unicode.IsUpper(firstChar) {
			result += str + " "
		}
	}
	return result, nil
}
