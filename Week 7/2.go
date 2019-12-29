package main

import "fmt"

type Album struct {
	name     string
	track    int
}

func Nexttrack(name string) *Album {
	GoC := new(Album)
	GoC.name = "Grimoire of Crimson"
	GoC.composer = "Team Grimoire"
	GoC.track = 1
}

func main() {
	fmt.Println("Music name :", GoC.name)
	fmt.Println("Composer :", GoC.composer)
	fmt.Println("Track :", GoC.track)
}
