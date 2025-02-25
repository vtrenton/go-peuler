package main

import "fmt"

func main() {
	i := 1
	for {
		triangle := calctriagle(i)
		factors := getfactors(triangle)
		fmt.Printf("%d: %v\n", triangle, factors)
		if len(factors) >= 500 {
			break
		}
		i++
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
