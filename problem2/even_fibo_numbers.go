package main

import "fmt"

func main() {
	fibslice := []int{0, 1}

	for range [4000000]struct{}{} {
		fib(&fibslice)
	}
	fmt.Println(fibslice)
}

func fib(fibslice *[]int) {
	s := *fibslice
	next := s[len(s)-2] + s[len(s)-1]
	*fibslice = append(s, next)
}
