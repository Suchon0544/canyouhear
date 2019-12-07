package main

import "fmt"

func main(){
	count := [5]string{"first" , "second" , "third" , "forth" , "fifth"}
	fmt.Println(count)
	Length := len(count)
	fmt.Println(Length)
	fmt.Println(count[1])
	//in "[]" is order that depend on index
}
