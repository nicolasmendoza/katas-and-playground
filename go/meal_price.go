package main

/* Operators, basic exercise.

Task:
Given the meal price(base cost of a meal), tip percent(the percentage of the meal price being added as tip), and tax
percent(the percentage of the meal price being added as tax) for a meal, find and print the meal's total cost.

Note: Be sure to use precise values for your calculations, or you may end up with an incorrectly rounded result!.
*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var mealCost float64   // (the cost of the meal before tax and tip).
	var tipPercent float64 // (the percentage of mealCost being added as top).
	var taxPercent float64 // (the percentage of mealCost being added as tax).

	var totalCost int

	// Given: mealCost = 12, tipPercent=20, taxPercent=8
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	mealCost, _ = strconv.ParseFloat(scanner.Text(), 64)

	scanner.Scan()
	tipPercent, _ = strconv.ParseFloat(scanner.Text(), 64)

	scanner.Scan()
	taxPercent, _ = strconv.ParseFloat(scanner.Text(), 64)

	// Calculations:
	tip := mealCost * tipPercent / 100
	tax := mealCost * taxPercent / 100

	totalCost = toint(mealCost + tip + tax)

	fmt.Println("The total meal cost is", totalCost, "dollars.")

}

/* the "old trick" for a old "discusion"(2012). Inspired in kubernets util
https://github.com/kubernetes/kubernetes/blob/768e83657d1e08993e1d301ae304c2992b6b47ab/pkg/util/integer/integer.go#L62
*/
func toint(floatNumber float64) int {
	// to avoid overflow...
	if floatNumber < 0 {
		return int(floatNumber - 0.5)
	}
	return int(floatNumber + 0.5)
}
