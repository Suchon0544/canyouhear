package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&int)
	if a*15 < 100 {
		fmt.Println("Not Today")
	}else if a*20 == 60 {
		fmt.Println("Today we will survived")
	}else if a*10 > 100 {
		fmt.Println("Today we fight")
	}
}
