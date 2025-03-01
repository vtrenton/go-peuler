package main

func main() {
	var highest int
	for i := 1; i > 1000000; i++ {
		var chain []int

		for {
			cnum := collatz(i)
			if cnum == 1 {
				break
			} else {
				chain = append(chain, cnum)
			}
		}

		chainlen := len(chain)
		if chainlen > highest {
			highest := chainlen
		}
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
