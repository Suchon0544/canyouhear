package main

import "fmt"

func iNeedHealing() {
	inh := recover()
	fmt.Print(inh)
}

func main() {
	defer iNeedHealing()
	panic("You didn't have any map.")
	var name map[int]string
	name[1] = "Outro : Wings"
	fmt.Println(name)
}
