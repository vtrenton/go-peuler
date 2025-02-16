package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func evendivisiblebyall(limit int) bool {
	for i := 1; i <= limit; i++ { // need to make sure we start at 1
		if num%i != 0 {
			return false
		}
	}
	return true
}
