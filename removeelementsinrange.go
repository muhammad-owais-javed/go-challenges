func RemoveElementsInRange(arr []float64, from, to int) []float64 {

	
  if from > to {
    from , to = to, from
  } 

  result := make([]float64, 0)
    
  for i := 0; i<len(arr) ; i++ {

    if i < from || i >= to {
      result = append(arr, i)
    }

  }

}