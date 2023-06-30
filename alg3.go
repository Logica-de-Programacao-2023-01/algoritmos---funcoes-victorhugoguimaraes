package main

import (
	"fmt"
	"strings"
)

func main() {
	strings := []string{"Hello", " ", "World!"}
	concatenated := concatenateStrings(strings)
	fmt.Println("Concatenação:", concatenated)
}

func concatenateStrings(strings []string) string {
	result := strings.Join(strings, "")
	return result
}
