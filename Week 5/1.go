package main

import "fmt"

func call_you_name(name string) func(string) string {
	return func(responce string) string {
		return name + responce
	}
}

func main() {
	fmt.Println("enter be like name time")
	fmt.Print("Plaese enter your name and what's time = ")
	var text string
	var time float32
	fmt.Scan(&text, &time)
	if time >= 18 {
		call_my_name := call_you_name(text)
		fmt.Println(call_my_name(" Good Night"))
		fmt.Print("It's time to go to bed")
	} else if time >= 15 {
		call_my_name := call_you_name(text)
		fmt.Println(call_my_name(" Good Evening"))
		fmt.Print("Come Back Home!")
	} else if time >= 12 {
		call_my_name := call_you_name(text)
		fmt.Println(call_my_name(" Good Afternoon"))
		fmt.Print("Have you had lunch, Right?")
	} else if time >= 6 {
		call_my_name := call_you_name(text)
		fmt.Println(call_my_name(" Good Morning"))
		fmt.Print("Have a good day! ,sunrise have come!")
	} else if time >= 1 {
		call_my_name := call_you_name(text)
		fmt.Println(call_my_name(" Good Night"))
		fmt.Print("It's time to go to bed")
	}
}
