package main

import (
	"fmt"
	"math"
)

func main() {
	factors := make([]int, 0)
	for i := 1; len(factors) <= 500; i++ {
		triangle := calctriagle(i)
		factors := getfactors(triangle)
		pfactors := primeFilter(factors)
		fmt.Printf("%d: %v\n", triangle, pfactors)
	}

}

func calctriagle(size int) int {
	var triange int
	for i := range size {
		triange += i
	}
	return triange
}

func getfactors(n int) []int {
	var factors []int
	limit := n / 2
	for i := 1; i <= limit; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

// lets see what trianges look like with only prime factors

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
