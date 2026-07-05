package sprint

func BalanceOut(arr []bool) []bool {
	var trueCounter int
	var falseCounter int
	for _, i := range arr {
		if i {
			trueCounter++
		} else  {
			falseCounter++
		}
		}
		
	if trueCounter > falseCounter {
		diff := trueCounter - falseCounter
		
		for i := 0; i < diff; i++ {
				arr = append(arr, false)
			}
		
		} else if falseCounter > trueCounter {
			diff := falseCounter - trueCounter
			
			for i := 0; i < diff; i++ {
				arr = append(arr, true)
			}
		}

		return arr
}
