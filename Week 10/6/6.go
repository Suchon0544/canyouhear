package main

import "fmt"

type number struct {
	prime int
	name  string
}

func main() {
	pn := number{
		prime: 2,
		name:  "Two",
	}
	fmt.Printf("First prime number : %v , (%v)", pn.prime, pn.name)
}
