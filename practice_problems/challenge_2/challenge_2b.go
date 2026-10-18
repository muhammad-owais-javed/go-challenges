package main

import (
	"fmt"
)

type BigInt []uint8

func main(){

	fmt.Println("Hello World")

	var num1 uint64 = 9956
	result := FromInt(num1)
	fmt.Printf("From Int to BigInt: (num1) = %v -> %v \n",num1, result)

	str := "00123"
	result = FromDecimal(str)
	fmt.Printf("From Decimal to BigInt: (str) = %v -> %v \n", str, result)

	a := FromDecimal("999")
	b := FromDecimal("1")
	sum := Add(a, b)

	fmt.Printf("sum: %d \n", sum)

	x := FromDecimal("3")
	y := FromDecimal("6")
	prod := Multiply(x, y)
	fmt.Printf("prod: %d \n", prod)

	a = FromInt(1)
	b = FromDecimal("1")
	for range 200 {
    	a, b = b, Add(a, b)
	}

	fmt.Println("ToDecimal:", ToDecimal(a)) 

}


func FromInt(i uint64) BigInt {

	if i == 0 {

		fmt.Println("i:", i)
		return BigInt{}
	
	}

	var b BigInt
	
	for i > 0 {
		b = append(b, uint8(i%10))
		i /= 10
	}
	
	return b
}

func FromDecimal(s string) BigInt {
	
	var parse []uint8

	for i := 0; i < len(s); i++ {
		char := s[i]
		if char >= '0' && char <= '9' {
			parse = append(parse, char - '0')
		}
	}

	start := 0
	
	for start < len(parse) && parse[start] == 0 {
		start ++
	}

	parse = parse[start:]

	if len(parse) == 0 {
		return BigInt{}
	}

	b := make(BigInt, len(parse))
	for i := 0; i < len(parse); i++{
		b[i] = parse[len(parse)-1-i]
	}

	return b
}

func ToDecimal(i BigInt) string {

	str := make([]byte, len(i))

	for j := 0; j < len(i); j++ {
		str[j] = i[len(i)-1-j]+'0'
	}

	return string(str)

}




func Add(a, b BigInt) BigInt {

	var result BigInt
	var carry uint8

	maxlen := len(a)
	if len(b) > maxlen {
		maxlen = len(b)
	}

	for i := 0; i < maxlen; i++ {

		var sum uint8 = carry

		if i < len(a) {
			sum += a[i]
		}

		if i < len(b) {
			sum += b[i]
		}

		result = append(result, sum % 10)
		carry = sum / 10

	}

	if carry > 0 {
		result = append(result, carry)
	}

	return result
}

func Multiply(a, b BigInt) BigInt {
	
	if len(a) == 0 || len(b) == 0 {
		return BigInt{}
	}

	// maximum number of digits in multiplication is len(a)+len(b)
	result := make(BigInt, len(a)+len(b))

	for i := 0; i < len(a); i++ {

		var carry uint8 = 0
		
		for j:= 0; j < len(b); j++ {
			
			multiply := (a[i] * b[j]) + carry + result[i+j]
			result [i+j] = multiply % 10
			carry = multiply / 10
		}

		result[i+len(b)] += carry
	}


	return result
}
