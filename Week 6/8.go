package main

import "fmt"

func main() {
	c := []int{}
	a := 2
	for a <= 100 {
		b := a * 12
		c = append(c, b)
		fmt.Println(c)
		a := a + 3
	}
}
