package main

import "fmt"

var choose [10]Chapter
var i int

type Chapter struct {
	RomOrder string
	name     string
}

func main() {
	choose[0] = Chapter{"I", "L1 : Loneliness"}
	choose[1] = Chapter{"II", "L2 : Liberation"}
	choose[2] = Chapter{"III", "L3 : Leverage"}
	choose[3] = Chapter{"IV", "L4 : Latent"}
	choose[4] = Chapter{"V", "L5 : Lapse"}
	choose[5] = Chapter{"VI", "L6 : Lunatic"}
	choose[6] = Chapter{"VII", "L7 : Legion"}
	choose[7] = Chapter{"VIII", "L8 : Lost"}
	choose[8] = Chapter{"IX", "L9 : Lament"}
	choose[9] = Chapter{"X", "L10 : Largo"}

	fmt.Println("Enter your number : ")
	fmt.Scan(&i)
	if i == 1 {
		fmt.Print(choose[0])
	} else if i == 2 {
		fmt.Print(choose[1])
	} else if i == 3 {
		fmt.Print(choose[2])
	} else if i == 4 {
		fmt.Print(choose[3])
	} else if i == 5 {
		fmt.Print(choose[4])
	} else if i == 6 {
		fmt.Print(choose[5])
	} else if i == 7 {
		fmt.Print(choose[6])
	} else if i == 8 {
		fmt.Print(choose[7])
	} else if i == 9 {
		fmt.Print(choose[8])
	} else if i == 10 {
		fmt.Print(choose[9])
	}
}
