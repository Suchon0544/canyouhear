package main

import "fmt"

type grimoireOfEmerald struct {
	name  string
	track int
}

func main() {
	var gOE [12]grimoireOfEmerald
	gOE[4] = grimoireOfEmerald{"Dainsleif ~The Gold Curse~", "5"}
	gOE[5] = grimoireOfEmerald{"Campo", "6"}

	fmt.Println("track name : ", gOE[4].track)
	fmt.Println("track name : ", gOE[5].name)
}
