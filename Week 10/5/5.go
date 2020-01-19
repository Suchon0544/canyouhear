package main

import "fmt"

type country struct {
	name        string
	people      string
	lang        string
	CapitalCity string
}

func main() {
	zh := country{
		name:        "China",
		people:      "Chinese",
		lang:        "Chinese",
		CapitalCity: "Beijing",
	}
	th := country{
		name:        "Thailand",
		people:      "Thai",
		lang:        "Thai",
		CapitalCity: "Bangkok",
	}
	fmt.Printf("%v infomation\n", zh.name)
	fmt.Printf("%v 's name : %v\n", zh.name, zh.people)
	fmt.Printf("\n")
	fmt.Printf("%v information\n", th.name)
	fmt.Printf("%v's Capital City : %v\n", th.name, th.CapitalCity)
}
