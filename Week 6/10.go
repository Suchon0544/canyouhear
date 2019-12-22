package main

import "fmt"

func main() {
	fmt.Println("Please enter your number = ")
	var a int
	fmt.Scan(&a)
	if a*5 < 20 {
		fmt.Println("Then no problem, I kill")
	} else if a*10 > 20 {
		fmt.Println("But, I dont care")
	}
	fmt.Println("You can control ma shyuu")
}
