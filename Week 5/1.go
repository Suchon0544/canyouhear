package main

import "fmt"

func call_you_name(name string) func(string) string {
	return func(responce string) string {
		return name + responce
	}
}

func main() {
	fmt.Println("enter be like name time(hours)")
	fmt.Print("Plaese enter your name and what's time = ")
	var text string
	var time int
	fmt.Scan(&text, &time)
	if time >= 18 {
		call_my_name := call_you_name(text)
		fmt.Println(call_my_name(" Good Night"))
	} else if time >= 15 {
		call_my_name := call_you_name(text)
		fmt.Println(call_my_name(" Good Evening"))
	} else if time >= 12 {
		call_my_name := call_you_name(text)
		fmt.Println(call_my_name(" Good Afternoon"))
	} else if time >= 6 {
		call_my_name := call_you_name(text)
		fmt.Println(call_my_name(" Good Morning"))
	} else if time >= 1 {
		call_my_name := call_you_name(text)
		fmt.Println(call_my_name(" Good Night"))
	}
}
//how lonely people do when they're alone....
