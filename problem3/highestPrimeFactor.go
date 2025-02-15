package main

import (
	"fmt"
	"math"
)

func main() {
	const base = 600851475143
	// get all the factors of n
	factors := getFactors(base)
	// loop through factors and determine which are prime
	primeFactors := primeFilter(factors)
	// print the highest
	fmt.Println(primeFactors[len(primeFactors)-1])

}

func getFactors(base int) []int {
	var factors []int
	limit := math.Sqrt(float64(base))
	for i := range int(limit) {
		if i == 0 {
			continue
		}
		if base%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

func primeFilter(factors []int) []int {
	var primeFactors []int
	for _, v := range factors {
		if isPrime(v) {
			primeFactors = append(primeFactors, v)
		}
	}
	return primeFactors
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
