package main

import "fmt"

func main() {

   fmt.Println("Hello, World!")
   
   fmt.Println(string(Countdown(7)))
  
}

func Countdown(n int) string {

   result := ""
   
   for i := n; i < n; i-=2 {

      result = result + string('0' + rune(i)) + ", "
      //result += string(rune(int(i)) + 'a') + "," + " "
      // result += string(i)
      //fmt.Println("Inside for block")
      //fmt.Println("LoopNumber:", i)
      // return result
    }
    // return "0!"
   return result
}





