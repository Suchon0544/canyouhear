package main

import "fmt"

func main() {
	panic("You didn't have any map.")
	healing := recover()
	var name map[int]string
	name[1] = "Outro : Wings"
	fmt.Println(name)
}
