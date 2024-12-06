package main

import (
	"fmt"
	"log"
)

func main() {
	var n int
	for {
		fmt.Println("Enter a positive integer greater than 0, or enter 0 to exit: ")
		// Scan(a ...interface{}) (n int, err error)
		_, err := fmt.Scan(&n)
		if err != nil {
			log.Print(err)
			continue
		}
		// Exit if "0" is entered
		if n == 0 {
			fmt.Println("Exiting...")
			break
		}
		// continue if a negative number is entered
		if n < 0 {
			continue
		}
		collatz(n)
	}
}

func collatz(n int) {
	var i int // counter variable
	fmt.Println("Number entered:", n)
	// 3n+1 if odd, n/2 if even
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		i++
		fmt.Println(n)
	}
	fmt.Println("This algorithm went through", i, "iterations.")
}
