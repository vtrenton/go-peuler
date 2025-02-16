package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var pal []int

	// need to do better than o(n^2)
	for i := 100; i <= 999; i++ {
		for j := i; j <= 999; j++ {
			product := i * j
			if ispalindrome(product) {
				pal = append(pal, product) //not appending in order
			}
		}
	}
	sort.Ints(pal)
	fmt.Println(pal[len(pal)-1])
}

func ispalindrome(num int) bool {
	strnum := strconv.Itoa(num)
	split := (len(strnum) + 2 - 1) / 2

	for i := range split {
		if strnum[i] != strnum[len(strnum)-1-i] {
			return false
		}
	}

	return true
}
