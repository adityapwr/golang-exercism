package accumulate

// Accumulate return a uppercase string of any input string
func Accumulate(inputString []string, operation func(string) string) (resultObj []string) {
	for _, i := range inputString {
		resultObj = append(resultObj, operation(i))
	}
	return resultObj
}
