package main

import (
	"fmt"
)

func main() {
  fmt.Println("Main Function")
  
 
  arr := [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}
  var num int = 9


  fmt.Println(arr[0][1], num)

  for i:=0 ; i < len(arr) ; i++ {
    for j:=0 ; j < len(arr) ; j++ {

      fmt.Println(arr[i][j])

    }


  }

}
