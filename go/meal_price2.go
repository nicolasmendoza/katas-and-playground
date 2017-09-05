package main

import (
	"fmt"
)

const (
	pi = 3.1416
)

func main(){
	var cost, tipPercent, taxPercent float64
	fmt.Scanln(&cost)
	fmt.Scanln(&tipPercent)
	fmt.Scanln(&taxPercent)
	fmt.Println(pi)
	tip:= cost * (tipPercent/100.0)
	tax := cost * (taxPercent/100.0)
	fmt.Printf("The total meal cost is %.0f dollars.", cost+tip+tax)
}

