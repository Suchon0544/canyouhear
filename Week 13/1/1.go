package main

import "fmt"

var pl [33]playlist

type playlist struct {
	name   string
	artist string
}

func allmusic(pl playlist) {
	fmt.Println(fmt.Sprintf("Name : %v  Artist : %v", pl.name, pl.artist))
}

func main() {
	pl[0] = playlist{name: "Adventure", artist: "Verdammt"}
	pl[1] = playlist{name: "Agust D", artist: "Agust D"}
	pl[2] = playlist{name: "Black Swan", artist: "BTS"}
	pl[3] = playlist{name: "Boys in Kaleidosphere", artist: "MILI"}
	pl[4] = playlist{name: "citanLu", artist: "ice"}
	pl[5] = playlist{name: "Crytal Snow", artist: "BTS"}
	pl[6] = playlist{name: "disappeared despair", artist: "Verdammt"}
	pl[7] = playlist{name: "Excalibur", artist: "MILI"}
	pl[8] = playlist{name: "Fantasy", artist: "JBJ"}
	pl[9] = playlist{name: "Flame Dark", artist: "KIVA"}
	pl[10] = playlist{name: "Grimoire of Crimson", artist: "Team Grimoire"}
	pl[11] = playlist{name: "Help me", artist: "Nu'est"}
	pl[12] = playlist{name: "Immoral", artist: "ice"}
	pl[13] = playlist{name: "I'll sleep when I am dead", artist: "Set It Off"}
	pl[14] = playlist{name: "J.B.J.(Intro)", artist: "JBJ"}
	pl[15] = playlist{name: "Kill the light", artist: "Set it Off"}
	pl[16] = playlist{name: "Lemonade", artist: "MILI"}
	pl[17] = playlist{name: "Moonchild", artist: "RM"}
	pl[18] = playlist{name: "New Rule", artist: "TXT"}
	pl[19] = playlist{name: "On my mind", artist: "JBJ"}
	pl[20] = playlist{name: "Overcome", artist: "Nu'est"}
	pl[21] = playlist{name: "Parodia Sonatina", artist: "ice"}
	pl[22] = playlist{name: "Qualia", artist: "KIVA"}
	pl[23] = playlist{name: "Restriction", artist: "Team Grimoire"}
	pl[24] = playlist{name: "Run Away", artist: "TXT"}
	pl[25] = playlist{name: "Seoul", artist: "RM"}
	pl[26] = playlist{name: "So far away", artist: "Agust D"}
	pl[27] = playlist{name: "Tokyo", artist: "RM"}
	pl[28] = playlist{name: "Touchin", artist: "Kang Daniel"}
	pl[29] = playlist{name: "Utopiosphere", artist: "MILI"}
	pl[30] = playlist{name: "Vitamins", artist: "MILI"}
	pl[31] = playlist{name: "What are you up to", artist: "Kang Daniel"}
	pl[32] = playlist{name: "Just be Star", artist: "JBJ"}
	p := allmusic(pl...)
	fmt.Println(p)
}
