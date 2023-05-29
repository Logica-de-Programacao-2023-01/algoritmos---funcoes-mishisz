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

func main() {
	numeros := []int{10, 5, 8, 12, 6}
	media, err := mediaAritmetica(numeros)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Média:", media)
	}
}
