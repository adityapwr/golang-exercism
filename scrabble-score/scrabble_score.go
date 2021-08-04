package scrabble

import (
	"unicode"
)

//Function Score take an input string and returns score on the basic of dictionary
func Score(word string) (score int) {
	group1 := []string{"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"}
	group2 := []string{"D", "G"}
	group3 := []string{"B", "C", "M", "P"}
	group4 := []string{"F", "H", "V", "W", "Y"}
	group5 := []string{"K"}
	group8 := []string{"J", "X"}
	group10 := []string{"Q", "Z"}
	for _, i := range word {
		_, found1 := find(group1, string(unicode.ToUpper(i)))
		if found1 {
			score = score + 1
		}
		_, found2 := find(group2, string(unicode.ToUpper(i)))
		if found2 {
			score = score + 2
		}
		_, found3 := find(group3, string(unicode.ToUpper(i)))
		if found3 {
			score = score + 3
		}
		_, found4 := find(group4, string(unicode.ToUpper(i)))
		if found4 {
			score = score + 4
		}
		_, found5 := find(group5, string(unicode.ToUpper(i)))
		if found5 {
			score = score + 5
		}
		_, found8 := find(group8, string(unicode.ToUpper(i)))
		if found8 {
			score = score + 8
		}
		_, found10 := find(group10, string(unicode.ToUpper(i)))
		if found10 {
			score = score + 10
		}
	}
	return
}

func find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
