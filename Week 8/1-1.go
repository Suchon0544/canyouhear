package main

import "fmt"

func minusThree(im int) int {
	if im%3 == 0 {
		fmt.Println(im)
	}
	return 0
}

func main() {
	minusThree(40)
}
