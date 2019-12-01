package myTruelyText

import 
	"fmt"
	"testing"

func TestShowText(t *testing.T){
	total := ShowText(10,5)
	if total = 15{
		fmt.Print("Finally")
	}
	if total != 15 {
		t.Error(total, "is not over")
	}
}