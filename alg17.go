package main

import (
	"fmt"
	"sort"
)

func main() {
	strings := []string{"banana", "abacaxi", "laranja"}
	sorted, err := sortStrings(strings)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Strings ordenadas:", sorted)
	}
}

func sortStrings(strings []string) ([]string, error) {
	if len(strings) == 0 {
		return nil, fmt.Errorf("slice vazio")
	}
	sorted := make([]string, len(strings))
	copy(sorted, strings)
	sort.Strings(sorted)
	return sorted, nil
}
