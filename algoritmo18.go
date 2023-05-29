package main

import "fmt"

func aplicarFuncaoEObterSoma(slice []int, fn func(int) int) (int, error) {
	if len(slice) == 0 {
		return 0, fmt.Errorf("O slice está vazio")
	}

	soma := 0
	for _, num := range slice {
		resultado := fn(num)
		soma += resultado
	}

	return soma, nil
}
