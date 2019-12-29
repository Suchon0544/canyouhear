package main

import "fmt"

type grimoireOfEmerald struct {
	name  string
	track int
}

func main() {
	var gOE [12]grimoireOfEmerald
	gOE[5] = grimoireOfEmerald{"Campo", "6"}

	fmt.Println(alb[8])
	fmt.Println("track name : ", gOE[5].name)
}
