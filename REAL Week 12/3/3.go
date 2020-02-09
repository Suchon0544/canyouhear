package main

import "fmt"

func main() {
	var dep int
	fmt.Print("Plaese enter your deposit = ")
	fmt.Scan(&dep)

	inter := (dep * 2) / 100

	balance := dep + inter

	fmt.Println(fmt.Sprintf("Your Deposit = %v", dep))
	fmt.Println(fmt.Sprintf("Your interst = %v", inter))
	fmt.Println(fmt.Sprintf("Your Balance = %v", balance))
}
