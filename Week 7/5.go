package main

import "fmt"

type Album struct {
	name  string
	track string
}

func (alb Album) listen() {
	fmt.Println("Have you ever hear this ,", alb.name)
}

type music struct {
	address string
	Album
}

func main() {
	var mus music
	mus.name = "Grimoire of Crimson"
	mus.listen()
}
