package main
import (
	"fmt"
	"reflect"
)

func main() {
  fmt.Println("Main Function")
  
  var min int = 1
  var max int = 5
  var size int = max - min + 1
  
  arr := make([]int, size)
  fmt.Println(reflect.TypeOf(arr))
  fmt.Println(len(arr))
  
  for i:= 0; i<size; i++ {
  	arr[i] = min+i

  }

  fmt.Println(arr)
    
}
