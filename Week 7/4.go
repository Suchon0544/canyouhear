package main

import "fmt"

type Album struct {
	name  string
	track string
}

func (alb Album) listen(){
	fmt.Println("Have you ever hear this ," , alb.name)
}

type music struct {
	address string
	alb		Album
}

func main() {
	first := Album{name : "Grimoire of Crimson"}
	mus := music{alb: first}
	mus.alb.listen()
}
