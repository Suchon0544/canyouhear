package main

import "fmt"

type Album struct {
	name  string
	track string
}

func main() {
	var alb [9]Album
	alb[0] = Album{"Grimoire of Crimson", "1"}
	alb[1] = Album{"Jormungand", "2"}
	alb[2] = Album{"G1ll35 d3 R415", "3"}
	alb[3] = Album{"upright", "4"}
	alb[4] = Album{"Vandetta", "5"}
	alb[5] = Album{"b4d dr4go0n", "6"}
	alb[6] = Album{"Nightmare System", "7"}
	alb[7] = Album{"Freedom", "8"}
	alb[8] = Album{"C18H27NO3", "9"}

	fmt.Println("track name : ", alb[8].name)
	fmt.Println("track name : ", alb[5].name)
}
