package main

import (
	"flag"
	"fmt"
	"github.com/goldennovember/control/pcag/wc"
	"os"
)

func main() {
	// Defining a boolean flag -l to count words instead of lines
	words := flag.Bool("w", false, "Count words")

	bytes := flag.Bool("b", false, "Count bytes")
	// Parsing the flags provided by the user
	flag.Parse()

	// Calling the count function to count the number of words (or lines)
	// received from the Standard Input and printing it out
	fmt.Println(wc.Count(os.Stdin, *words, *bytes))
}
