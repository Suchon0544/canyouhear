package main

import "fmt"

func devide(a, b int) int {
	return a / b
}

func helpMe() {
	helpYou := recover()
	if helpYou == 0 {
		fmt.Println("You result should not be zero")
		main()
	}
}

func main() {
	defer helpMe()
	c := devide(15, 32)
	fmt.Println(c)
}
