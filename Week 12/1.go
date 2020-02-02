package main

import (
	"fmt"
	"math/rand"
)

func desk() []int {
	dsk := make([]int, 0)
	a := rand.Intn(13)
	b := rand.Intn(13)
	c := rand.Intn(13)
	d := rand.Intn(13)
	e := rand.Intn(13)

	dsk = append(dsk, a, b, c, d, e)
	return dsk

}

func main() {
	card := desk()
	fmt.Println(card)
}
