package main

import "fmt"

func main() {
	//magic_num := 1000
	//magic_limit := int(math.Sqrt(float64(magic_num)))
	// sqrt(1000) ~= 31.6

	// calc permutations of all possible combinations of 3 numbers that equal n
	// each of these numbers need to have an even sqrtrt with a remainder of 0

	// RULE: the difference between all squared numbers is an incrementing odd number appart
	// 1,3,5,7,9
	// our 3 numbers will be an odd distance appart

	// RULE: a + b + c = 1000; b - a == n; c - b == n + 2

	// the goal is 3 numbers next to each other that == 1000

	//var result int
	var prod int
	a := 1
	b := a + 1
	c := b + 1

	for {
		prod = a + b + c
		fmt.Println(prod)
		if prod == 999 {
			fmt.Printf("a: %d, b: %d, c: %d\n", a, b, c)
			fmt.Printf("%d\n", square(a)*square(b)*square(c))
			break
		} else {
			a++
			b++
			c++
		}
	}
}

func square(num int) int {
	return num * num
}
