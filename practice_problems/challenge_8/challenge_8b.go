package main

import (
	"fmt"
	"errors"
)

const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"


func main(){

	encoded := EncodeBase62(4567)
	decodedInt, decodedErr := DecodeBase62("WA10")

	fmt.Printf("\nEncode: %s, Decode: %d, Err: %v\n", encoded, decodedInt, decodedErr)

}


func EncodeBase62(n uint64) string {
   
	var result []byte

	if n == 0 {
		return "0"
	}

	for n > 0 {

		remainder := n%62

		char := alphabet[remainder]

		result = append(result, char)

		n /= 62

	}

	for i := 0; i < len(result)-1 ; i ++ {

		for j := len(result)-1; j > i ; j-- {
			result[i], result[j] = result[j], result[i]
		}

	}
	
	return string(result)
}

func DecodeBase62(s string) (uint64, error) {
   
	var total uint64

	if s == "0" {
		return 0, errors.New("empty String")
	}
	
	for i := 0; i < len(s) ; i++ {

		char := s[i]

		val := -1

		for j := 0; j < len(alphabet); j++ {
			if alphabet[j] == char {
				val = j
				break
			}
		}
	
		if val == -1 {
			return 0, errors.New("invalid character in string")
		}

		total = (total * 62) + uint64(val)


	}

	return total, nil

}