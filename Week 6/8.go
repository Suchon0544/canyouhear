package main

import "fmt"

func main() {
	a := 1
	for {
		a := a + 3
		b := a + 12
		fmt.Println(b)
		if a <= 20 {
			break
		}
	}
}
