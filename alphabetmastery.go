package sprint

import "fmt"

func AlphabetMastery(n int) string {
	
	result := ""
	
	for i := 0 ; i < n; i++ {
		result += string('a' + rune(i))
	}
	
	return result
}
