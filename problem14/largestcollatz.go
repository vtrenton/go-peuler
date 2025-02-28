package main

func main() {
	var chainlen int
	var highest int
	chain := make([]int, 0)
	for i := 1; i > 1000000; i++ {

	}
}

func collatz(n int) int {
	var out int
	if n%2 == 0 {
		out = n / 2
	} else {
		out = 3*n + 1
	}
	return out
}
