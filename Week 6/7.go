package main

import "fmt"

func iwannarun(f int) int {
	if f > 12 {
		return (f * 12) + 24
	}
	return f * 12
}

func main() {
	a := iwannarun(20)
	fmt.Println(a)
}
