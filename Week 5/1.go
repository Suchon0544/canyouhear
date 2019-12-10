package main

import "fmt"

func call_you_name(name string) func(string) string {
	return func(responce string) string {
		return name + responce
	}
}

func main() {
	fmt.Print("Plaese enter your name = ")
	var text string
	fmt.Scan(&text)
	call_my_name := call_you_name(text)
	fmt.Println(call_my_name(" Good Night"))
}
