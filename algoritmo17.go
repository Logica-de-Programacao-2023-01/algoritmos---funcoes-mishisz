package main

import (
	"fmt"
	"sort"
	"strings"
)

func ordenarStrings(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", fmt.Errorf("O slice está vazio")
	}

	sort.Strings(slice)

	novaString := strings.Join(slice, ", ")

	return novaString, nil
}
