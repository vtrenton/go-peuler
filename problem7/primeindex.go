package main

import (
	"fmt"
	"math"
)

func main() {
	var counter int
	num := 2
	for {
		if isprime(num) {
			counter++
		}
		if counter == 10001 {
			break
		}
		num++
	}
	fmt.Println(num)
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
