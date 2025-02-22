package main

import (
	"fmt"
)

func main() {
	fmt.Println("vim-go")
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
	for i := range limit {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}
