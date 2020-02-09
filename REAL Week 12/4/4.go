package main

import "fmt"

type Country struct {
	name        string
	lang        string
	capitalcity string
}

func correcting(co *Country) {
	co.lang = "British English"
}

func main() {
	var country = Country{
		name:        "United Kingdom of Great Britain and Northern Ireland | ",
		lang:        "English | ",
		capitalcity: "London",
	}

	correcting(&country)

	fmt.Println(country)
}
