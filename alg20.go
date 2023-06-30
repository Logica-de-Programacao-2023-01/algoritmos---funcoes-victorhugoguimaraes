package main

import (
	"fmt"
)

func main() {
	strings := []string{"Hello", "world", "Go"}
	longStrings, err := filterLongStrings(strings)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Strings longas:", longStrings)
	}
}

func filterLongStrings(strings []string) ([]string, error) {
	if len(strings) == 0 {
		return nil, fmt.Errorf("slice vazio")
	}
	longStrings := []string{}
	for _, str := range strings {
		if len(str) > 5 {
			longStrings = append(longStrings, str)
		}
	}
	return longStrings, nil
}
