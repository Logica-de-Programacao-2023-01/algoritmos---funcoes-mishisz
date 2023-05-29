package main

import (
	"errors"
	"fmt"
)

func aplicarFuncaoEmCadaElemento(nums []int, funcao func(int) int) ([]int, error) {
	if len(nums) == 0 {
		return nil, errors.New("Slice está vazio")
	}

	resultados := make([]int, len(nums))
	for i, num := range nums {
		resultados[i] = funcao(num)
	}

	return resultados, nil
}

func dobrarNumero(num int) int {
	return num * 2
}
