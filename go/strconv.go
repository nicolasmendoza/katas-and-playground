package main

/*
Playing with strconv, bufio, and stdin

From stdin/Input receives three values:

12
4.0
 with strconv, bufio, and stdin


*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var i uint64 = 8
	var d float64 = 4.0
	var s string = "Playing"

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	// base , bitSize
	i2, _ := strconv.ParseUint(scanner.Text(), 10, 64)

	scanner.Scan()
	d2, _ := strconv.ParseFloat(scanner.Text(), 64)

	scanner.Scan()
	s2 := scanner.Text()

	fmt.Println(i + i2)

	// fmt, prec, bitSize
	fmt.Println(strconv.FormatFloat(d+d2, 'f', '1', 64))

	fmt.Println(s + s2)

	/*
		Output:

		20
		9.0
		Playing with strconv, bufio, and stdin
	*/
}
