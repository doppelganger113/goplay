package main

import (
	"log"
	"os"
	_ "github.com/doppelganger113/goplay/examples/network_example/matchers"
	"github.com/doppelganger113/goplay/examples/network_example/search"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("president")
}
