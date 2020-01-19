package main

import "fmt"

func main() {
	var withd int
	fmt.Println("ระบุจำนวนเงิน :")
	fmt.Scan(&withd)

	if withd <= 99 {
		fmt.Println("ไม่สามารถทำรายการได้")
	}
	if withd >= 100 {
		H := withd / 100
		if H >= 5 {
			I := H / 5
			fmt.Printf("ธนบัตร 500 บาท จำนวน %v ใบ", I)
		}
		fmt.Printf("ธนบัตร 100 บาท จำนวน %v ใบ", H)
	}
}
