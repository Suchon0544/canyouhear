package main

import "fmt"

type songs struct {
	name string
	composer string
	track int
}

func (music songs) listen() {
	fmt.Println("Have you ever hear this ," music.name)
}

type playlist struct {
	library string
	songs
}

func main(){
	playlist.name = "MIC Drop"
	playlist.listen()
}