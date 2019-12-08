package main

import "fmt"

func main(){
	counts := [2][2][2]string{{{"Fisrt","Second"} , {"Eins","Zwei"}} , {{"Hana","Tul"},{"Ichi","Ni"}}}
	fmt.Println(counts)
	fmt.Println(counts[1][1][1])
}
