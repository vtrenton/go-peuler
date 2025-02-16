package main

func main() {
	var pal []int
	for i := 100; i <= 999; i++ { //what the fuck
		if isPrime() {
			for j := 100; j <= 999; j++ {
				if isPrime() {
					pal = append(pal, j)
				}
			}
		}
	}
}
