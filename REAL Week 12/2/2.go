package main

import "fmt"

func deposit() {
	var dep float64
	inr := 2.125 / 100
	fmt.Scan(&dep)
	inter := dep * inr
	return inter
}

func main() {
	de := deposit
	fmt.Print(de)
}
