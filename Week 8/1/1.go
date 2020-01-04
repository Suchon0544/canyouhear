//ฟังก์่นที่รับพารามิเตอร์ 2 ตัวเป็นตัวลข หากเลขที่รับาเป็นจำนวนติดลบให้คืนค่าเป็น error

package main

import (
	"fmt"
)

type Minus interface {
	showError()
}

type Insert struct {
	input int
}

func (in Insert) showError() {
	fmt.Println(in.input)
}

func main() {
	var m Minus
	m = Insert{"irressible"}
	m.showError()
}
