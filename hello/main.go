package main

import "fmt"

func main() {

 reverse := ReverseAlphabetValue('b')
 fmt.Println("Hello, World!")
 fmt.Println(reverse, string(reverse))
}

func ReverseAlphabetValue(ch rune) rune {

 hold := 'z' - ch
 return hold + 'a'

}


