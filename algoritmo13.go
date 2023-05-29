package main

import (
	"errors"
	"fmt"
	"strconv"
)

func somaDigitos(numero int) (int, error) {
	if numero < 0 {
		return 0, errors.New("Número negativo não é válido")
	}

	strNumero := strconv.Itoa(numero)
	soma := 0

	for _, char := range strNumero {
		digito, _ := strconv.Atoi(string(char))
		soma += digito
	}

	return soma, nil
}
