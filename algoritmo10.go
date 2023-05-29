package main

import (
	"errors"
	"fmt"
	"sort"
)

func ordenarValores(nums []int) ([]int, error) {
	if len(nums) == 0 {
		return nil, errors.New("Slice está vazio")
	}

	copia := make([]int, len(nums))
	copy(copia, nums)

	sort.Ints(copia)

	return copia, nil
}
