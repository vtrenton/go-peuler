package main

import (
	"fmt"
	"math"
)

func main() {
	maxno := 2000000
	var sum int

	for i := range maxno {
		if isprime(i) {
			sum += i
		}
	}
	fmt.Println(sum)
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
