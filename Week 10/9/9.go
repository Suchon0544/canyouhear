package main

import (
	"fmt"
	"time"
)

type country struct {
	name        string
	people      string
	lang        string
	CapitalCity string
}

func main() {
	start := time.Now()
	zh := country{
		name:        "China",
		people:      "Chinese",
		lang:        "Chinese",
		CapitalCity: "Beijing",
	}
	fmt.Printf("%v infomation\n", zh.name)
	fmt.Printf("%v 's name : %v\n", zh.name, zh.people)
	fmt.Print(time.Since(start))

	//why they're 0s ????
}
