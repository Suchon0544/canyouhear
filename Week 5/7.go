package main

import "fmt"

func plus() func() int {
	beg := 0
	return func() int {
		beg := beg + 2
		return beg
	}
}

func main() {
	nextplus := plus()
	fmt.Println(nextplus())
}
