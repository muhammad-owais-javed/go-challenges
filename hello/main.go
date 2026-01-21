package main

import "fmt"

func main() {

 //reverse := ReverseAlphabetValue('b')
 fmt.Println("Hello, World!")
 //fmt.Println(reverse, string(reverse))
 var i int = 5
 if (i > 0 ) {
  fmt.Println("i is positive")
 
 }
 
}

func ReverseAlphabetValue(ch rune) rune {

 hold := 'z' - ch
 return hold + 'a'

}


