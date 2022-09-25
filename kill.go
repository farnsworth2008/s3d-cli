package main

import (
	"log"
)

// Kill anything that is not `nil`
//
// We clearly do not need to "kill" it so we simply return in that case. We are
// Nihilists.
func kill(err interface{}) {
	if err != nil {
		log.Fatal(err)
		killOn(err != nil, "err")
	}
}

// Kill on a true condition
func killOn(condition bool, because string) {
	if condition {
		panic(because)
	}
}
