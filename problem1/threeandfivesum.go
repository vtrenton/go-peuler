package main

import "fmt"

func main() {
	var sum int
	for i := range 1000 {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
