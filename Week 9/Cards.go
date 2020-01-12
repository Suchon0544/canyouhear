package main

import (
	"fmt"
	"math/rand"
)

var ran []cards = make([]cards, 100)
var code []int = make([]int, 100)
var draw [53]cards
var i int
var n int

type cards struct {
	order string
	sign  string
}

func Check(k int) int {

	for j := 0; j <= n; j++ {
		if code[j] == k {
			fmt.Println(ran[j])
			return 1
		}
	}
	return 0
}

func main() {
	draw[1] = cards{"A", "Spade"}
	draw[2] = cards{"2", "Spade"}
	draw[3] = cards{"3", "Spade"}
	draw[4] = cards{"4", "Spade"}
	draw[5] = cards{"5", "Spade"}
	draw[6] = cards{"6", "Spade"}
	draw[7] = cards{"7", "Spade"}
	draw[8] = cards{"8", "Spade"}
	draw[9] = cards{"9", "Spade"}
	draw[10] = cards{"10", "Spade"}
	draw[11] = cards{"J", "Spade"}
	draw[12] = cards{"Q", "Spade"}
	draw[13] = cards{"K", "Spade"}
	draw[14] = cards{"A", "Heart"}
	draw[15] = cards{"2", "Heart"}
	draw[16] = cards{"3", "Heart"}
	draw[17] = cards{"4", "Heart"}
	draw[18] = cards{"5", "Heart"}
	draw[19] = cards{"6", "Heart"}
	draw[20] = cards{"7", "Heart"}
	draw[21] = cards{"8", "Heart"}
	draw[22] = cards{"9", "Heart"}
	draw[23] = cards{"10", "Heart"}
	draw[24] = cards{"J", "Heart"}
	draw[25] = cards{"Q", "Heart"}
	draw[26] = cards{"K", "Heart"}
	draw[27] = cards{"A", "Diamond"}
	draw[28] = cards{"2", "Diamond"}
	draw[29] = cards{"3", "Diamond"}
	draw[30] = cards{"4", "Diamond"}
	draw[31] = cards{"5", "Diamond"}
	draw[32] = cards{"6", "Diamond"}
	draw[33] = cards{"7", "Diamond"}
	draw[34] = cards{"8", "Diamond"}
	draw[35] = cards{"9", "Diamond"}
	draw[36] = cards{"10", "Diamond"}
	draw[37] = cards{"J", "Diamond"}
	draw[38] = cards{"Q", "Diamond"}
	draw[39] = cards{"K", "Diamond"}
	draw[40] = cards{"A", "Club"}
	draw[41] = cards{"2", "Club"}
	draw[42] = cards{"3", "Club"}
	draw[43] = cards{"4", "Club"}
	draw[44] = cards{"5", "Club"}
	draw[45] = cards{"6", "Club"}
	draw[46] = cards{"7", "Club"}
	draw[47] = cards{"8", "Club"}
	draw[48] = cards{"9", "Club"}
	draw[49] = cards{"10", "Club"}
	draw[50] = cards{"J", "Club"}
	draw[51] = cards{"Q", "Club"}
	draw[52] = cards{"K", "Club"}
redo:
	fmt.Print("input your number : ")
	fmt.Scan(&n)
	if n != 0 {
		if Check(i) == 0 {
			var ran1 = draw[rand.Intn(53)]
			ran[n] = ran1
			code[n] = 1
			n++
			fmt.Println(ran1)
		}
		goto redo
	}
}
