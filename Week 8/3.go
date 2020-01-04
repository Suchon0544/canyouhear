package main

import "fmt"

func main() {
	for im := 1; im < 1001; im++ {
		if im%3 == 0 {
			fmt.Println("fizz", im)
		} else if im%5 == 0 {
			fmt.Println("Buzz", im)
		}
	}
}
