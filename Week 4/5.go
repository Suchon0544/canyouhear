package main

import "fmt"

func main(){
	counts := [2][2][2]string
	{
		{"first" , "second"} ,
		{"Eins" , "Zwei"},
	} , 
	{
		{"Hana" , "Tul"} ,
		{"Ichi" , "Ni"},
	}
	fmt.Println(counts[1][0][0])
}