package main

import (
	"fmt"
	"math"
)

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func getPrimes(n int) ([]int, error) {
	if n < 0 {
		return nil, fmt.Errorf("O número deve ser não negativo")
	}

	primes := make([]int, 0)

	for i := 2; i <= n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	return primes, nil
}
