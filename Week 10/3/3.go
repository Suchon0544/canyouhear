package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	for i := 1; i <= 3000; i++ {
		file, err := os.Create(fmt.Sprintf("%v.txt", i))
		if err != nil {
			return
		}
		for j := 1; j <= 1000; j++ {
			defer file.Close()
			file.WriteString("Suchonlaphat\r\nSuwanaphokin\r\n22\r\n")
		}
	}
	fmt.Println(time.Since(start))
}
