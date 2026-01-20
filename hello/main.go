package main

import "fmt"

func main() {
    
 a=Casting(3.456)
 fmt.Println("Hello, World!")
 fmt.Print(a)
}

func Casting(n float64) int {
  var x int = int(math.Round(n))
  return x

}


