package main

import "fmt"

func main() {
	limit := 4000000
	fibo := fibogen(limit)
	fibeven := fibeven(fibo)

	fmt.Println(sumfibeven(fibeven))

}

func fibogen(limit int) []int {
	var fibo []int
	var last int
	i := 1

	for i <= limit {
		fibo = append(fibo, i)
		fiblen := len(fibo)
		if fiblen >= 2 {
			last = fibo[fiblen-2]
		} else {
			last = 0
		}
		i += last
	}
	return fibo
}

func fibeven(fibo []int) []int {
	var fibeven []int
	for _, v := range fibo {
		if v%2 == 0 {
			fibeven = append(fibeven, v)
		}
	}
	return fibeven
}

func sumfibeven(fibeven []int) int {
	var sum int
	for _, v := range fibeven {
		sum += v
	}
	return sum
}
