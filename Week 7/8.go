package main

import "fmt"

type knightVol2 struct {
	name  string
	track string
}

func main() {
	var kV [2]knightVol2
	kV[0] = knightVol2{"Silent Oath", "1"}
	kV[1] = knightVol2{"Fight for Judge", "2"}

	fmt.Println("track name : ", alb[8].name)
	fmt.Println("track name : ", alb[5].name)
}
