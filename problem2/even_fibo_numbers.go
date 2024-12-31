package main

import "fmt"

func main() {
	fibslice := []int{0, 1}

	for range [4000000]struct{}{} {
		fib(&fibslice)
	}
	fmt.Println(fibslice)
}

func fib(fibslice *[]int) []int {
	localfib := *fibslice
	next := localfib[len(localfib)-2] + localfib[len(localfib)-1]
	localfib = append(localfib, next)
	return localfib
}
