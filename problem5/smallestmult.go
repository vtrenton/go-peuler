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
	var primeexp []int
	for _, prime := range primes {
		exponent := 1
		for {
			// look forward and see if we can interate without a bust!
			next := int(math.Pow(float64(prime), float64(exponent+1)))
			if next <= limit {
				exponent++
			} else {
				break
			}
		}
		primeexp = append(primeexp, exponent)
	}

	return primeexp
}

func getlcm(primes, exponents []int) int {
	result := 1
	var multiples []int

	for i := range len(primes) {
		pow := int(math.Pow(float64(primes[i]), float64(exponents[i])))
		multiples = append(multiples, pow)
	}

	for _, multiple := range multiples {
		result *= multiple
	}

	return result
}
