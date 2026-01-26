package main

func GenerateRange(min, max int) []int{

  if ( max < min ) {
	var arr []int
	return arr
  }
	
  var size int = max - min + 1 
  
  arr := make([]int, size)
  
  for i := 0; i < size; i++ {
  	arr[i] = min+i
  }

  return arr

}
