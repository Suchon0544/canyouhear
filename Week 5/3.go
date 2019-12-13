package main

import "fmt"

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}

func main() {
	fac := factorial(3)
	fmt.Println(fac)
}

// they work like real factorial((n!-1) *n)
//so it's how they work at the first should be 3*2
// next (3*2)*1 (number == 0 now)
//that why fac(3) == 6

//for me that similar for(loop) they work almost same way ,but diiferece called
