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

func geteven(list []int) []int {
	var evenlist []int
	for _, v := range list {
		if v%2 == 0 {
			evenlist = append(evenlist, v)
		}
	}
	return evenlist
}

func sumslice(fibeven []int) int {
	var sum int
	for _, v := range fibeven {
		sum += v
	}
	return sum
}
