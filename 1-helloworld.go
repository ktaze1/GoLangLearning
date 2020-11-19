package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	loops()
}

func loops() {
	// some type of loops
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	} // classical for loop

	for 3 > 2 { //condition here
		// while loop
	}

	for {
		// some task
		// infinite loop
	}
}

func rangeInLoop() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
