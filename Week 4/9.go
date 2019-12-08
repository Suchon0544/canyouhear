package main

import "fmt"

func main(){
	orders := [] string {"Eins","Zwei","Drei","Viert"}
	orders2 := [] string {"Wu","Liu","Qi","Ba"}
	fmt.Println(orders)
	fmt.Println(orders2)

	copy(orders , orders2)
	fmt.Println(orders)
	fmt.Println(orders2)
}
