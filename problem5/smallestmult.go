package main

import (
	"fmt"
	"math"
)

func main() {
	const limit = 20
	primes := getprimestolimit(limit)
	exponents := getexponents(primes, limit)
	lcm := getlcm(primes, exponents)

	fmt.Println(lcm)
}

func getprimestolimit(limit int) []int {
	var primes []int
	for i := range limit {
		if isprime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func isprime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func getexponents(primes []int, limit int) []int {
	var exponent int
	var primeexp []int
	for prime := range primes {
		for exponent = 1; int(math.Pow(float64(prime), float64(exponent))) < limit; exponent++ {
		}
		primeexp = append(primeexp, exponent)
	}
	return primeexp
}

func getlcm(primes, exponents []int) int {
	result := 1
	var multiples []int
	for prime := range primes {
		for exp := range exponents {
			mult := int(math.Pow(float64(prime), float64(exp)))
			multiples = append(multiples, mult)
		}
	}

	for multiple := range multiples {
		result *= multiple
	}

	return result
}
