package main

import "fmt"

type dep struct {
	dep float64
}

func deposit(dep *Dep) float64 {
	fmt.Print("Plaese enter your deposit = ")
	fmt.Scan(&dep)
	return dep
}

func interest(dep *Dep) float64 {
	inr := 2.125 / 100
	inter := dep * inr
	return inter
}

func depositing(mon float64) float64 {
	fmt.Scan(&mon)
	depo := dep
	total := depo + mon
	return total

}

func main() {
	de := deposit(1)
	fmt.Print(fmt.Sprintf("Your interest = %v Bath", de))
}
