package sprint

import "fmt"

func AlphabetMastery(n int) string {

	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	var result string

	if n > 0 && n < 26 {
		for ( i:=0; i<n ; i++ ){
			result += string(alphabet[i])
		}
	}
	
	return result

}
