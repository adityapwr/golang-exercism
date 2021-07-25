// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap takes year as a param and returns bool true if year is a leap year
package leap

// IsLeapYear return true if year passes is a leap year else returns false
func IsLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}
