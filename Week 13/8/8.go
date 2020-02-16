package main

import (
	"fmt"
	"math/rand"
)

type playlist struct {
	name   string
	artist string
}

func main() {
	a := playlist{name: "Overcome", artist: "Nu'est"}
	b := playlist{name: "Parodia Sonatina", artist: "ice"}
	c := playlist{name: "Qualia", artist: "KIVA"}
	d := playlist{name: "Restriction", artist: "Team Grimoire"}
	e := playlist{name: "Run Away", artist: "TXT"}
	f := playlist{name: "Seoul", artist: "RM"}
	g := playlist{name: "So far away", artist: "Agust D"}
	h := playlist{name: "Tokyo", artist: "RM"}
	i := playlist{name: "Touchin", artist: "Kang Daniel"}
	j := playlist{name: "Utopiosphere", artist: "MILI"}
	k := playlist{name: "Vitamins", artist: "MILI"}
	l := playlist{name: "What are you up to", artist: "Kang Daniel"}
	m := playlist{name: "Just be Star", artist: "JBJ"}

	fmt.Print("Playing next track from ")
	x := rand.Intn(14)
	fmt.Println(x)

	fmt.Println("Now Playing")
	if x == 13 {
		fmt.Println(fmt.Sprintf("Name : %v Artist : %v", a.name, a.artist))
	} else if x == 1 {
		fmt.Println(fmt.Sprintf("Name : %v Artist : %v", b.name, b.artist))
	} else if x == 2 {
		fmt.Println(fmt.Sprintf("Name : %v Artist : %v", c.name, c.artist))
	} else if x == 3 {
		fmt.Println(fmt.Sprintf("Name : %v Artist : %v", d.name, d.artist))
	} else if x == 4 {
		fmt.Println(fmt.Sprintf("Name : %v Artist : %v", e.name, e.artist))
	} else if x == 5 {
		fmt.Println(fmt.Sprintf("Name : %v Artist : %v", f.name, f.artist))
	} else if x == 6 {
		fmt.Println(fmt.Sprintf("Name : %v Artist : %v", g.name, g.artist))
	} else if x == 7 {
		fmt.Println(fmt.Sprintf("Name : %v Artist : %v", h.name, h.artist))
	} else if x == 8 {
		fmt.Println(fmt.Sprintf("Name : %v Artist : %v", i.name, i.artist))
	} else if x == 9 {
		fmt.Println(fmt.Sprintf("Name : %v Artist : %v", j.name, j.artist))
	} else if x == 10 {
		fmt.Println(fmt.Sprintf("Name : %v Artist : %v", k.name, k.artist))
	} else if x == 11 {
		fmt.Println(fmt.Sprintf("Name : %v Artist : %v", l.name, l.artist))
	} else if x == 12 {
		fmt.Println(fmt.Sprintf("Name : %v Artist : %v", m.name, m.artist))
	}

}
