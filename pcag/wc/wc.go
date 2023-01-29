package wc

import (
	"bufio"
	"io"
)

func Count(r io.Reader, countWords bool, countBytes bool) int {

	// A scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)

	// If the count words or bytes flag is set
	// the scanner will split type to words or bytes (default is split by lines)
	if countWords {
		scanner.Split(bufio.ScanWords)
	} else if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	// Defining a counter
	wc := 0

	// For every word or line scanned, add 1 to the counter
	for scanner.Scan() {
		wc++
	}
	// Return the total
	return wc
}
