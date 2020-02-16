package main

import "fmt"

func main(){
	package main

import "fmt"

type playlist struct {
	name   string
	artist string
}

func main() {
	a := playlist{name: "Adventure", artist: "Verdammt"}
	b := playlist{name: "Agust D", artist: "Agust D"}
	c := playlist{name: "Black Swan", artist: "BTS"}
	d := playlist{name: "Boys in Kaleidosphere", artist: "MILI"}
	e := playlist{name: "citanLu", artist: "ice"}
	f := playlist{name: "Crytal Snow", artist: "BTS"}
	g := playlist{name: "disappeared despair", artist: "Verdammt"}
	h := playlist{name: "Excalibur", artist: "MILI"}
	i := playlist{name: "Fantasy", artist: "JBJ"}
	j := playlist{name: "Flame Dark", artist: "KIVA"}
	k := playlist{name: "Grimoire of Crimson", artist: "Team Grimoire"}
	l := playlist{name: "Help me", artist: "Nu'est"}
	m := playlist{name: "Immoral", artist: "ice"}
	n := playlist{name: "I'll sleep when I am dead", artist: "Set It Off"}
	o := playlist{name: "J.B.J.(Intro)", artist: "JBJ"}
	p := playlist{name: "Kill the light", artist: "Set it Off"}
	q := playlist{name: "Lemonade", artist: "MILI"}
	r := playlist{name: "Moonchild", artist: "RM"}
	s := playlist{name: "New Rule", artist: "TXT"}
	t := playlist{name: "On my mind", artist: "JBJ"}

	var x int
	fmt.Print("Choose your music : ")
	fmt.Scan(&x)

	fmt.Println("Now Playing")
	if x == 1 {
		fmt.Println(fmt.Sprintf("Name : %v Artist : %v", a.name, a.artist))
	} else if x == 2 {
		fmt.Println(fmt.Sprintf("Name : %v Artist : %v", b.name, b.artist))
	} else if x == 3 {
		fmt.Println(fmt.Sprintf("Name : %v Artist : %v", c.name, c.artist))
	} else if x == 4 {
		fmt.Println(fmt.Sprintf("Name : %v Artist : %v", d.name, d.artist))
	} else if x == 5 {
		fmt.Println(fmt.Sprintf("Name : %v Artist : %v", e.name, e.artist))
	} else if x == 6 {
		fmt.Println(fmt.Sprintf("Name : %v Artist : %v", f.name, f.artist))
	} else if x == 7 {
		fmt.Println(fmt.Sprintf("Name : %v Artist : %v", g.name, g.artist))
	} else if x == 8 {
		fmt.Println(fmt.Sprintf("Name : %v Artist : %v", h.name, h.artist))
	} else if x == 9 {
		fmt.Println(fmt.Sprintf("Name : %v Artist : %v", i.name, i.artist))
	} else if x == 10 {
		fmt.Println(fmt.Sprintf("Name : %v Artist : %v", j.name, j.artist))
	} else if x == 11 {
		fmt.Println(fmt.Sprintf("Name : %v Artist : %v", k.name, k.artist))
	} else if x == 12 {
		fmt.Println(fmt.Sprintf("Name : %v Artist : %v", l.name, l.artist))
	} else if x == 13 {
		fmt.Println(fmt.Sprintf("Name : %v Artist : %v", m.name, m.artist))
	} else if x == 14 {
		fmt.Println(fmt.Sprintf("Name : %v Artist : %v", n.name, n.artist))
	} else if x == 15 {
		fmt.Println(fmt.Sprintf("Name : %v Artist : %v", o.name, o.artist))
	} else if x == 16 {
		fmt.Println(fmt.Sprintf("Name : %v Artist : %v", p.name, p.artist))
	} else if x == 17 {
		fmt.Println(fmt.Sprintf("Name : %v Artist : %v", q.name, q.artist))
	} else if x == 18 {
		fmt.Println(fmt.Sprintf("Name : %v Artist : %v", r.name, r.artist))
	} else if x == 19 {
		fmt.Println(fmt.Sprintf("Name : %v Artist : %v", s.name, s.artist))
	} else if x == 20 {
		fmt.Println(fmt.Sprintf("Name : %v Artist : %v", t.name, t.artist))
	}
}

}