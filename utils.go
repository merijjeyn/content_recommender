package main

import "fmt"

// Just a layer of abstraction over logging so that we can easily upgrade it/mess with it, etc.
func log(message string) {
	fmt.Print(message, "\n")
}
