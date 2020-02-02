package main

import (
	"fmt"
	"math/rand"
)

func main() {
	bts := make([]int, 4)

	b := rand.Intn(13)
	bts = append(bts, b)
	fmt.Print(bts)
}
