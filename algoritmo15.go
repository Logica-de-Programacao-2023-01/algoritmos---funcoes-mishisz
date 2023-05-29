package main

import "fmt"

func numeroPresente(numero int, slice []int) (bool, error) {
	if len(slice) == 0 {
		return false, fmt.Errorf("O slice está vazio")
	}

	for _, num := range slice {
		if num == numero {
			return true, nil
		}
	}

	return false, nil
}
