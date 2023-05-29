package main

import (
	"errors"
	"fmt"
	"sort"
)

func segundoMaiorValor(nums []int) (int, error) {
	if len(nums) < 2 {
		return 0, errors.New("Slice deve conter pelo menos dois elementos")
	}

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	segundoMaior := nums[1]
	return segundoMaior, nil
}
