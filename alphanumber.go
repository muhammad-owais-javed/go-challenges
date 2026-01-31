package main

// import "fmt"

func AlphaNumber(n int) string {

	negative := ""
	if n == 0 {
		return "a"
	}
	if n < 0 {
		negative += "-"
		n = n * -1
	}
	result := ""
	for n > 0 {
		digit := n % 10
		char := string(rune('a' + digit))
		result = char + result
		n = n / 10
	}
	return negative + result
}
