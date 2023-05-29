package main

import (
	"fmt"
)

func encontrarPosicao(nums []int, valor int) int {
	for i, num := range nums {
		if num == valor {
			return i
		}
	}
	return -1
}
