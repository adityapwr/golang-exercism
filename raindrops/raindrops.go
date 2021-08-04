//Package raindrops converts ints to raindrop codes
package raindrops

import "fmt"

// Convert take an input integer and returns string raindrops. Capital case Convert for exporting function outside package
func Convert(input int) string {
	resultString := ""
	//Checking for factor of 3
	if input%3 == 0 {
		resultString += "Pling"
	}
	//Checking for factor of 5
	if input%5 == 0 {
		resultString += "Plang"
	}
	//Checking for factor of 7
	if input%7 == 0 {
		resultString += "Plong"
	}
	if resultString == "" {
		return fmt.Sprint(input)
	}
	return resultString
}
