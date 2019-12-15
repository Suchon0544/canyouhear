package main

import "fmt"

func listen(songs string) func(string) string {
	return func(composer string) string {
		return songs + composer
	}
}

func main() {
	var name string
	var compos string
	fmt.Scan(&name, &compos)
	you := listen(name + "_" + compos)
	fmt.Println(you(" Good song"))
}
