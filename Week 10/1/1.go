package main

import (
	"fmt"
)

func main() {
	var char string
	fmt.Scan(&char)
	if char == "A" {
		fmt.Println("A Cure for Wellness")
		fmt.Println("After Shock 1976")
		fmt.Println("Agora")
		fmt.Println("Alien Convenant")
	} else if char == "D" {
		fmt.Println("Dunkirk")
	} else if char == "G" {
		fmt.Println("Gladiator")
	} else if char == "I" {
		fmt.Println("Inception")
	} else if char == "K" {
		fmt.Println("Killing EVE")
		fmt.Println("Kill Your darling")
	} else if char == "M" {
		fmt.Println("Mary Queen Scots")
	}
}
