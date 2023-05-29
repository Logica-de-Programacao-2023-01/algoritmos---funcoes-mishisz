package main

import (
	"fmt"
	"strings"
)

func contarVogais(str string) int {
	count := 0
	vogais := "aeiouAEIOU"

	for _, char := range str {
		if strings.ContainsRune(vogais, char) {
			count++
		}
	}

	return count
}
