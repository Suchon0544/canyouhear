package main

import "fmt"

func listen(songs string) func(string) string {
	return func(composer string) string {
		return songs + composer
	}
}

func main() {
	you := listen("Restriction")
	fmt.Println(you(" Team grimoire"))
}
