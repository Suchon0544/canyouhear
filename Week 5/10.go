package main

import "fmt"

func count(numOrd int) func(int) int {
	return func(numAra int) int {
		return numOrd + numAra
	}
}

func main() {
	var num1 int
	fmt.Scan(&num1)
	cnt := count(num1)
	fmt.Println(cnt(num1))
}
