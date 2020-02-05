package main

import (
	"fmt"
	"math/rand"
)

type Hand struct {
	Cards  []int
	Couple int
}

func Deal() {
	a := rand.Intn(13)
	b := rand.Intn(13)
	c := rand.Intn(13)
	d := rand.Intn(13)
	e := rand.Intn(13)

	Cards = append(Cards, a, b, c, d, e)

}
func main() {
	fmt.Println()
}
