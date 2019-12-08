package main

import "fmt"

func main(){
	orders := [] string {"Eins","Zwei","Drei","Viert"}
	orders = append(orders ,"Wu","Liu","Qi","Ba")
	fmt.Println(orders)
	deleteIndex := 3
	orders = append(orders[deleteIndex+1:])
	fmt.Println(orders)
}
