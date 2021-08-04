// Package hamming calculate the Hamming Distance between two DNA strands.
package hamming

import (
	"errors"
)

// Distance accepts two input string and return hamming distance
func Distance(a, b string) (int, error) {
	ra := []rune(a)
	rb := []rune(b)
	if len(ra) != len(rb) {
		return 0, errors.New("Not Equal")
	}
	count := 0
	for i := range ra {
		if ra[i] != rb[i] {
			count++
		}
	}
	return count, nil
}
