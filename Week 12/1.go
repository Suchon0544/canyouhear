package main

import (
	"fmt"
	"math/rand"
)

func desk() []int {
	Dsk := make([]int, 0)
	a := rand.Intn(13)
	b := rand.Intn(13)
	c := rand.Intn(13)
	d := rand.Intn(13)
	e := rand.Intn(13)

	Dsk = append(Dsk, a, b, c, d, e)
	return Dsk

}

func newcard() {
	f := rand.Intn(13)
	Dsk = append(Dsk, f)
	return Dsk

}

func main() {
	card := desk()
	fmt.Println(card)
}
