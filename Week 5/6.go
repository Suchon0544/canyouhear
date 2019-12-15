package main

import "fmt"

func tried() {
	fmt.Println("Vitamins")
}

func main() {
	defer tried()
	fmt.Println("With a billoin worldful of <3")
}