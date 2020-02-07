package main

import "fmt"

func deposit(dep float64) float64 {
	fmt.Print("Plaese enter your deposit = ")
	fmt.Scan(&dep)
	inr := 2.125 / 100
	inter := dep * inr
	return inter
}

func main() {
	de := deposit(1)
	fmt.Print(fmt.Sprintf("Your interest = %v Bath", de))
}
