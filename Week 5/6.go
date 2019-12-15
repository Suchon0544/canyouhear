package main

import "fmt"

func tried() {
	fmt.Println("Vitamins")
}

func tolive(){
	fmt.Println("Extension of you")
}

func main() {
	defer tried()
	defer tolive()
	fmt.Println("With a billoin worldful of <3")
}

// defer == defered func will active before running func
// order(s) upside-down
