package main

import (
	"fmt"
	"math"
	"os"
	"text/tabwriter"
)

func main() {
	// func NewWriter(output io.Writer, minwidth, tabwidth, padding int, padchar byte, flags uint) *Writer
	// tabwriter.Debug flag adds cell borders
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)

	// sets initial conditions
	a := 0
	b := 1
	var c int
	var ratio float64

	phi := 1.61803398874989484

	// manually print initial conditions
	fmt.Fprintln(w, "i\tFib(i)\tFib(i)/Fib(i-1)\tAbsolute Error")
	fmt.Fprintln(w, "0", "\t", "0", "\t", "\t")
	fmt.Fprintln(w, "1", "\t", "1", "\t", "\t")

	// the fibonacci algorithm starting at index 2
	for i := 2; i < 41; i++ {
		c = a + b
		ratio = float64(c) / float64(b)
		fmt.Fprintln(w, i, "\t", c, "\t", ratio, "\t", math.Abs(ratio-phi))
		a = b
		b = c

	}
	// Flush should be called after the last call to Write to ensure that any data buffered in the Writer is written to output.
	// Any incomplete escape sequence at the end is considered complete for formatting purposes.
	w.Flush()

	fmt.Println("\nThe Golden Ratio is denoted by the Greek letter Phi φ")
	fmt.Println("φ to 15 places is: ", phi) // 15 places seems to be the cap
	fmt.Println("It is expressed in its algebraic form as: [1 + sqrt(5)] / 2")
	fmt.Println("as the solution to the quadratic equation: x^2 - x - 1 = 0")
}
