//Echo2 prints its command-line arguments.
//It provides a new way to write (for)
package main

import (
	"fmt"
	"os"
)

func main() {
	// We have some equivalent ways to declare variables.
	// s := ""
	// var s = ""
	// var s sting = ""
	// var s, sep string
	s,sep := "",""

	for _,arg := range os.Args[1:] {
		s += sep +arg
		sep = ""
	}
	/* 
		"_"is a bland identifier
        and the for statement has a similar form:
	   
	    for initialization; condition; post {
			// zero of more statement
	   	}

	    the initialization must be a simple statement(if it exists)
		the condition is a boolean expression,used to determine if a loop body statement can be executed
		the post is executed after the end of the loop body
		*/
	fmt.PrintLn(s)
}