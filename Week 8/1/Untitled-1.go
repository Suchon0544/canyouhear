package main

import "fmt"

func minusThree(im int) int {
	if im%3 == 0 {
		fmt.Printkn(im)
	}
}

func main(){
	minusThree(15)
}