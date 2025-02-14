package main

import "fmt"

func main() {
	var last int
	for i := 1; i <= 4000000; i += last {
		fmt.Println(last)
		fmt.Println(i)
		last = i
	}
}
