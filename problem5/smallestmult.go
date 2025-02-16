package main

import (
	"fmt"
	"math"
)

func main() {
	const limit = 20
	num := limit
	var i int
	for {
		i++
		fmt.Printf("i is %d\n", i)
		if evendivbyall(num, limit) {
			fmt.Println(num)
			break
		}
		num = int(math.Pow(limit, float64(i)))
		fmt.Printf("num is %d\n", num)
	}
	fmt.Println(num)
}

func evendivbyall(num, limit int) bool {
	for i := 1; i <= limit; i++ { // need to make sure we start at 1
		if num%i != 0 {
			return false
		}
	}
	return true
}
