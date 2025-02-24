package main

func main() {

}

// Triangle numbers are in a patter of diff between each number is plus one
/*
0
1 = 1
3 = 2
6 = 3
10
15
21
*/

func calctriagle(size int) int {
	var triangle int
	for diff := range size {
		triangle += diff
	}
}

//func calctriagle(size int) int {
//	var triange int
//	for i := range size {
//		triange += i
//	}
//	return triange
//}

func getfactors(n int) []int {
	var factors []int
	limit := n / 2
	for i := range limit {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}
