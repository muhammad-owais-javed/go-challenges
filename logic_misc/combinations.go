package main

import "fmt"

func Combinations() string {
 result := ""
	
  for i := 0; i <10; i++ {
    for j := i; j <10; j++ {
	   for  k := j; k < 10; k++ {
	     if i != j && k != j {
	 	    result += fmt.Sprintf("%d%d%d, ", i, j, k)
		    }
	    }
    }
  }

  result=result[:len(result)-2]
  return result

}
