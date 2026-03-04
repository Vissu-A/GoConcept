package goroutiensexp

import (
	"fmt"
)

func CheckStatusRoutine() {
	fmt.Println("Checking web status concurrently.")

	weblinks := []string{
		"http://google.com",
		"http://facebook.com",
		"http://twitter.com",
		"http://linkedin.com",
	}

	for idx, link := range weblinks {
		fmt.Println("web link index: ", idx, "----", "web link: ", link)
		go CheckStatus(link) // creating go routine by prefixing the function call with "go" keyword.

	}

}

// we can define a goroutine by simple prefixing with "go" keyword while calling the function.
// Syntax:
// go <calling-the-function>

// These go routines like "async event loop" in fast api.
// if one go routine blicked by some I/O operation like network call, then execution will be swiched to other go routine by Go scheduler.
// it uses one cpu core and one thread but switches between routines like event loop in fast api.
// this is concurency.

// By default Go will create a "Main Routine" when we execute the program.
// if the program contains any go routine functions and blocking the execution by I/O, then it will create a child go routie to run with other part of the program.

// Bug in GoRoutines
//------------------
// Main program/ GoRoutine will exit once it's work done, it will not wait until all child go routines are completed.
