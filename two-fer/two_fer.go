//Package twofer excepts a name and return string "One for x, one for me",
// if no name is provided it would return "One for you, one for me"
package twofer

// ShareWith accepts a variable string and returns a string
func ShareWith(name string) string {
	if name != "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
