package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func stringsComLetraMaiuscula(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("Slice está vazio")
	}

	var result []string

	for _, str := range slice {
		if len(str) > 0 && unicode.IsUpper(rune(str[0])) {
			result = append(result, str)
		}
	}

	return strings.Join(result, ", "), nil
}
