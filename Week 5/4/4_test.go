package andfunc

import "testing"

func TestAnd(t *testing.T) {
	r := and(true, false)
	if r != false {
		t.Error("Your Logics went erong")
	}
}
