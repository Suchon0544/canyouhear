package find

import "testing"

func TestGetDrive(t *testing.T) {
	got := getDrive()
	want := []string
	if got != want {
		t.Errorf("got %v , want %v" , got , want)
	}
}
