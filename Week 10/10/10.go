package main

import (
	"fmt"
	"time"
)

func call_you_name(name string) func(string) string {
	return func(responce string) string {
		return name + responce
	}
}

func main() {
	start := time.Now()
	fmt.Println("enter be like name time")
	fmt.Print("Plaese enter your name and what's time = ")
	var text string
	var tim float32
	fmt.Scan(&text, &tim)
	if tim >= 18 {
		call_my_name := call_you_name(text)
		fmt.Println(call_my_name(" Good Night"))
		fmt.Print("It's time to go to bed")
	} else if tim >= 15 {
		call_my_name := call_you_name(text)
		fmt.Println(call_my_name(" Good Evening"))
		fmt.Print("Come Back Home!")
	} else if tim >= 12 {
		call_my_name := call_you_name(text)
		fmt.Println(call_my_name(" Good Afternoon"))
		fmt.Print("Have you had lunch, Right?")
	} else if tim >= 6 {
		call_my_name := call_you_name(text)
		fmt.Println(call_my_name(" Good Morning"))
		fmt.Print("Have a good day! ,sunrise have come!")
	} else if tim >= 1 {
		call_my_name := call_you_name(text)
		fmt.Println(call_my_name(" Good Night"))
		fmt.Print("It's time to go to bed")
	}
	fmt.Println(time.Since(start))
}
