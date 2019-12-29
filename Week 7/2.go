package main

import "fmt"

type Album struct {
	name  string
	track int
}

func Nexttrack(name string) *Album {
	GoC := new(Album)
	GoC.name = "Grimoire of Crimson"
	GoC.track = 1
	return &GoC
}

func main() {
	fmt.Println(Album{"Music name :", GoC.name})
	fmt.Println(Album{"Track :", GoC.track})
}
