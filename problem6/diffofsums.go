package main

import "fmt"

func main() {
	var sum int
	var sum_sq int
	for i := 1; i <= 100; i++ {
		sum += i
		sum_sq += square(i)
	}

	diff := square(sum) - sum_sq
	fmt.Println(diff)
}

func square(num int) int {
	return num * num
}
