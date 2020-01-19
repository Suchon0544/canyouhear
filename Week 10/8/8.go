package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 1; i <= 3000; i++ {
		fmt.Println("all right")
	}
	fmt.Print(time.Since(start))
}
