package main

import "fmt"

func main() {
	fibslice := []int{0, 1}

	for i_ := range 4000000 {
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
