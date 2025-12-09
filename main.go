/**
 * @author Miles Aube
 * @version 1.0.0
 * @date 2025-12-07
 * @fileoverview This program calculates base exponent using a While loop.
 */

package main

import ("fmt")

func main() {
	// variables
	var base int
	var exponent int
	var result int = 1
	var counter int = 0

	// input from the user
	fmt.Print("Enter the base: ")
	fmt.Scan(&base)

	fmt.Print("Enter the exponent: ")
	fmt.Scan(&exponent)

	// calculate using a while loop (Go uses for as while)
	for counter < exponent {
		result = result * base
		counter = counter + 1
	}

	// output the answer
	fmt.Printf("\n%d raised to the power of %d is: %d\n", base, exponent, result)
}