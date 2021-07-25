// Package Hamming calculate the Hamming Distance between two DNA strands.
package hamming

import (
	"errors"
)

// Distance accepts two input string and return hamming distance
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Not Equal")
	} else {
		count := 0
		for i := range a {
			if a[i] != b[i] {
				count += 1
			}
		}
		return count, nil
	}
}
