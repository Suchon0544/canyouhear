package myTruelyText

import "testing"

func TestShowText(t *testing.T){
	total := ShowText(3,5)
	
	if total != 15 {
		t.Error(total, "is not over")
	}
}