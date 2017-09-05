/*
This "package" solve the basic next task:

Task:
Given an integer, n, perform the following conditional actions:

if n is odd, print Weird
if n is even and in the inclusive range of 2 to 5, print Not Weird
if n is even and in the inclusive range of 6 to 20, print Weird
if n is even and greater than 20, print Not Weird

Complete the stub code provided in your editor to print whether or not n
is weird.

Input Format.

A single line containing a positive integer, n.

Constrains:

1 <= n <= 100

Output Format:
Print Weird if the number is weird; otherwise, print Not Weird.

Sample input 0
3
Sample output 0
Weird
Sample input 1
24
Sample output 1
Not Weird
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isEven(n uint64) bool {

	if n == 0 {
		return false
	}

	if n%2 == 0 {
		return true
	}

	return false
}

func isBetween(i, min, max uint64) bool {

	if i >= min && i <= max {
		return true
	}

	return false
}

func main() {

	var n uint64

	weird := "Weird"
	notWeird := "Not Weird"

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	n, _ = strconv.ParseUint(scanner.Text(), 10, 64)
	var nIsEven bool = isEven(n)

	switch {

	case !nIsEven:
		fmt.Println(weird)

	case nIsEven && isBetween(n, 2, 5):
		fmt.Println(notWeird)

	case nIsEven && isBetween(n, 6, 20):
		fmt.Println(weird)

	case nIsEven && n > 20:
		fmt.Println(notWeird)

	default:
		fmt.Println(notWeird)
	}

}
