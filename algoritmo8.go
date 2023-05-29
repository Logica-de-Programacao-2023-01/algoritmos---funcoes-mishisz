package main

import (
	"errors"
	"fmt"
)

// Função que retorna um novo slice com apenas os números pares do slice de entrada
// Caso o slice esteja vazio, retorna um erro
func obterNumerosPares(nums []int) ([]int, error) {
	if len(nums) == 0 {
		return nil, errors.New("Slice está vazio")
	}

	pares := []int{}
	for _, num := range nums {
		if num%2 == 0 {
			pares = append(pares, num)
		}
	}

	return pares, nil
}
