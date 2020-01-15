package doomovies

import "testing"

func TestDoo(t *testing.T) {
	got := doo("I")
	want := "Inception"
	if got != want {
		t.Error("Not Found")
	}
}
