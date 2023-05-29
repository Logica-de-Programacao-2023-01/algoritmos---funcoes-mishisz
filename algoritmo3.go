package main

import (
	"fmt"
	"strings"
)

// Função que retorna a concatenação de todas as strings de um slice
func concatenarStrings(stringsSlice []string) string {
	concatenado := strings.Join(stringsSlice, "")
	return concatenado
}
