package main

import "fmt"

type Album struct {
	name  string
	track string
}

func main() {
	var alb [9]Album
	alb[8] = Album{"C18H27NO3", "9"}

	fmt.Println(alb[8])
	fmt.Println(alb[8].name)
}
