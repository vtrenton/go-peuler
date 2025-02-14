package main

import "fmt"

func main() {
	var last int
	i := 1
	for i <= 4000000 {
		last = i
		fmt.Println(last)
		fmt.Println(i)
		i += last
	}
}
