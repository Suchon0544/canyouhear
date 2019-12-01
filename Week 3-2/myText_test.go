package myText

import "testing"

func TestText(t *testing.T){
	e := Text(2,3)

	if e != 5 {
		t.Error("You go wrong, Dude.")
	}
}