//Package proverb returns a rhyme given an array of words
package proverb

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) (rhymeOutput []string) {
	if len(rhyme) == 0 {
		return
	}
	// for indx, i := range rhyme {
	// 	if indx == len(rhyme)-1 {
	// 		rhymeOutput = append(rhymeOutput, "And all for the want of a "+rhyme[0]+".")
	// 	} else {
	// 		rhymeOutput = append(rhymeOutput, "For want of a "+i+" the "+rhyme[indx+1]+" was lost.")
	// 	}
	// }
	// return
	len := len(rhyme) - 1
	for i := 0; i < len; i++ {
		rhymeOutput = append(rhymeOutput, "For want of a "+rhyme[i]+" the "+rhyme[i+1]+" was lost.")
	}
	rhymeOutput = append(rhymeOutput, "And all for the want of a "+rhyme[0]+".")
	return
}

// BenchmarkProverb-8   	  604770	      1879 ns/op	    1264 B/op	      30 allocs/op
// PASS
// ok  	proverb	2.274s
