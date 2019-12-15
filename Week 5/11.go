package main

import "fmt"

//ผลบวกของเลขทั้งหมด
func sum(number ...float64) float64 {
	total := 0.00
	for _, n := range number {
		total = total + n
	}
	return total
}

//ค่าเฉลี่ย
func average(number ...float64) float64 {
	ans1 := 0.00
	ans2 := 0.00
	ind := i
	for i, n := range number {
		ans1 = ans1 + n

	}

	ans2 = ans1 / (ind + 1)
	return ans2
}

func main() {
	x := sum(1, 2, 3, 4, 5)
	y := average(1, 2, 3, 4, 5)
	fmt.Println(x)
	fmt.Println(y)
}

//ค่าเฉลี่ย
//ค่าที่มากที่สุด
//ค่าที่น้อยที่สุด
