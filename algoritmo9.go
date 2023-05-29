package main

import (
	"errors"
	"fmt"
	"strings"
)

// Função que retorna um novo slice com todas as palavras da string de entrada
// Considera que as palavras são separadas por espaços em branco
// Caso a string seja vazia, retorna um erro
func obterPalavrasDaString(str string) ([]string, error) {
	if len(str) == 0 {
		return nil, errors.New("String está vazia")
	}

	palavras := strings.Split(str, " ")

	return palavras, nil
}
