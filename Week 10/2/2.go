package main

import "os"

func main() {
	file, err := os.Create("sth.txt")
	if err != nil {
		return
		defer file.Close()
	}
	file.WriteString("Suchonlaphat \n")
	file.WriteString("Suwanaphokin \n")
	file.WriteString("22")
}
