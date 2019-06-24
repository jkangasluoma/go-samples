package main

import (
	"fmt"
	"strconv"
)

// A runnable program always has a main function.
func main() {
	fmt.Println(youGo())
	ControlFlow()
}

// Private func in lower case:
func youGo() string {
	return "Go You!"
}

// Sometimes even code comments are required by convention, see below:

// ControlFlow Showcases control flow structures. Note upper case - public visibility, documentation required.
func ControlFlow() {
	// Sequence: from top to bottom:

	flow := "Sequential flow"
	fmt.Print(flow)

	flow = " from top to bottom."
	fmt.Println(flow)

	// Strict typing:
	const number = 42
	flow = strconv.Itoa(number)
	fmt.Printf("%T %+v\n", flow, flow)

	// Loop:

	for i := 0; i < 100; i++ {
		fmt.Printf("%d,", i)
	}

	a := 1
	const b = 3

	for a < b {
		fmt.Print("\nA is nearing b.")
		a++
	}

	fmt.Println("")

	// Conditional:

}
