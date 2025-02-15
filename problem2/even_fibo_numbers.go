package main

import "fmt"

func main() {
	limit := 4000000
	fibo := fibogen(limit)
	fibeven := fibeven(fibo)

	fmt.Println(fibeven[len(fibeven)-1])

}

func fibogen(limit int) []int {
	var fibo []int
	var last int
	i := 1

	for i <= limit {
		fibo = append(fibo, i)
		if len(fibo) >= 2 {
			last = fibo[len(fibo)-2]
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
