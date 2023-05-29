package main

import (
	"fmt"
	"strings"
)

func substituirVogais(texto string) (string, error) {
	if texto == "" {
		return "", fmt.Errorf("A string está vazia")
	}

	vogais := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}

	for _, vogal := range vogais {
		texto = strings.ReplaceAll(texto, vogal, "*")
	}

	return texto, nil
}
