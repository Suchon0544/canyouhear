package main

import "fmt"

func main(){
	orders := [] string {"Eins","Zwei","Drei","Viert"}
	orders = append(orders ,"Wu","Liu","Qi","Ba")
	fmt.Println(orders)
	orders_by_num := []string{"one","two","three","four","five","six","seven","eight"}
	orders = append(orders_by_num)
	fmt.Println(orders)
}

//slice by order not index [count from left and slice : count from right and slice]
//append(var name that you want to add sth , sth that you wanna add to)