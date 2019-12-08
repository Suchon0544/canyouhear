package main

import "fmt"

func main(){
	orders := [] string {"Eins","Zwei","Drei","Viert"}
	orders2 := [] string {"Wu","Liu","Qi","Ba"}
	fmt.Println(orders)
	fmt.Println(orders2)
	fmt.Println("---------------------")

	copy(orders , orders2)
	fmt.Println(orders)
	fmt.Println(orders2)
	fmt.Println("---------------------")

	orders3 := [] string {"Kyu","Juu","Juu-Ichi","Juu-Ni"}
	copy(orders2 , orders3)
	fmt.Println(orders)
	fmt.Println(orders2)
	fmt.Println(orders3)
	fmt.Println("---------------------")
}
