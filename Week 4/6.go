package main

import "fmt"

func main(){
	orders := [] string {"Eins","Zwei","Drei","Viert"}
	fmt.Println(orders[3:])
}

//slice by order not index [count from left and slice : count from right and slice]