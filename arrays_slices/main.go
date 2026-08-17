package main

import (
	"fmt"
)

func main(){

	s := make([]int, 3)
	arr := make([]int, 3, 5)

	fmt.Printf("Slice: %d\n", s)
	fmt.Printf("Array: %d\n", arr)

}