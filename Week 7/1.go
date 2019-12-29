package main

import "fmt"

type Album struct {
	name     string
	composer string
	track    int
}

func main() {
	GoC := new(Album)
	GoC.name = "Grimoire of Crimson"
	GoC.composer = "Team Grimoire"
	GoC.track = 1

	fmt.Println("Music name :", GoC.name)
	fmt.Println("Composer :", GoC.composer)
	fmt.Println("Track :", GoC.track)
}
