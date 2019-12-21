package main

import "fmt"

func main() {
	name := make(map[int]string)
	name[1] = "Intro : Boys Meet Evil"
	fmt.Println(name)

	name[2] = "Outro : Wings"
	name := map[int]string{3:"Intro : Serendipity" , 4:"Outro : Her"}
	fmt.Println(name)
}
