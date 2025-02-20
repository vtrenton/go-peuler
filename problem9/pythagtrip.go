package main

import (
	"fmt"
)

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

	// if a and c are odd -> the result will be even

	//var result int
	nums := [3]int{1, 2, 3}

	for {
		var sum int

		fmt.Println(nums)

		for _, v := range nums {
			sum += v
		}
		fmt.Println(sum)
		if sum >= 1000 {
			var squares [3]int
			sum = 0
			for i, v := range nums {
				squares[i] = square(v)
				sum += squares[i]
			}
			fmt.Println(sum)
			break
		} else {
			for i, v := range nums {
				nums[i] = v + 2
			}
		}
	}
}

func square(num int) int {
	return num * num
}
