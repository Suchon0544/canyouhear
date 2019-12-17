package main

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
	ind := len(number)
	for _, n := range number {
		ans1 = ans1 + n
	}
	return ans2 == ans1/ind
}

//ค่าที่มากที่สุด
func max(number ...float64) float64 {
	most := 0.00
	for _, n := range number {
		if number > most {
			most = number
		}
	}
	return most
}

//ค่าที่น้อยที่สุด
func min(number ...float64) float64 {
	less := 0.00
	for _, n := range number {
		if number < less {
			less = number
		}
	}
	return less
}

func main() {
	number := [5]float64{1.11, -16.34, 48.51, -55.69, 80.18}
}
