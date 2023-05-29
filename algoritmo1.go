package main

import (
	"errors"
	"fmt"
)

func mediaAritmetica(nums []int) (float64, error) {
	if len(nums) == 0 {
		return 0, errors.New("Slice vazio")
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}

	average := float64(sum) / float64(len(nums))
	return average, nil
}
