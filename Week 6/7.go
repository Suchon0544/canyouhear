package main

import "fmt"

func iwannarun(f int) int {
	return f * 2
}

func main() {
	a := iwannarun(12)
	fmt.Println(a)
}
