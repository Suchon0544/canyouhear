package everheardthis

func TestHeard(t *tessting.T) {
	input := "V"
	want := "Singularity"
	got := heard(text)
	if got != want {
		t.Errorf("We dont have this on playlist")
	}
}
