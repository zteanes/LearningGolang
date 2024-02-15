package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

// outlines how for loops works in Go
func forLoop() {
	fmt.Println("\nFOR LOOPS -------------")
	sum := 0
	// it's literally java loops
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// init and post statements are optional
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// there is no while loop in Go, there's just for
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// can also just say for to have an infinite loop
	/*
		for {
			fmt.Println("INFINITE!")
		}
	*/
	fmt.Println("END FOR -----------------")
}

// outlines how the if statement works
func ifStatements(x, n, lim float64) float64 {
	fmt.Println("\nIF STATEMENTS ------------")
	y := 10.123
	square := 1.234
	if y >= 10 {
		square = math.Sqrt(y)
	}
	fmt.Println(square)

	// can do something right before if statement and enter based on result
	if v := math.Pow(x, n); v < lim { // only return whenever v is less than the limit
		return v
	} else {
		fmt.Println(v) // v also accessible in the else block
	}
	fmt.Println("END IF ------------")
	return square
}

// this function outlines how switch statements work
func switchStatements() {
	fmt.Println("\nSWITCH STATEMENT ------------")
	// same as other languages except no break, only executes the one select
	fmt.Println("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Println(os)
	}

	// evaluates from top to bottom and just stops after one succeeds
	// also works with no condition
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good night!")
	}

	fmt.Println("END SWITCH ------------")
}

func deferStatement() {
	fmt.Println("\nDEFER FUNCTION ------------")
	defer fmt.Println("END DEFER ------------") // see what I did there!
	// defers the execution of a function until the rest returns
	// evaluates it immediately but not executed until
	defer fmt.Println("World!")
	fmt.Println("Hello")

	// deferred statements are pushed onto stack and returned in LIFO order
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

func main() {
	forLoop()
	ifStatements(1.23, 2.34, 100.00)
	switchStatements()
	deferStatement()
}
