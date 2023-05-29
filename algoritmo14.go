package main

import "fmt"

func numerosComuns(slice1, slice2 []int) ([]int, error) {
	if len(slice1) == 0 || len(slice2) == 0 {
		return nil, fmt.Errorf("Um dos slices está vazio")
	}

	numeros := make([]int, 0)

	for _, num1 := range slice1 {
		for _, num2 := range slice2 {
			if num1 == num2 {
				numeros = append(numeros, num1)
				break
			}
		}
	}

	return numeros, nil
}
