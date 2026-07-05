package main

import "fmt"

func Pairs() string {

   result := ""
    for i := 0; i<100; i++ {
	for j := i; j<100; j++ {
	   if i != j {
	     result += fmt.Sprintf("%02d %02d, ", i, j)
		}
	}
  }
  result = result[:len(result)-2]
//  fmt.Println("Result:", string(result))
  return result
}
