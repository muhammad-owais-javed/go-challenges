package sprint

func BetweenLimits(from, to rune) string {

	if from > to {
		from, to = to, from
	}

	result := []rune{}

	for r := from + 1; r < to; r++ {
		result = append(result, r)
	}

	return string(result)
}
/* 
if from > to {
	from, to = to, from
 }

 var resultrune []rune
 for i:=from+1 ; i<to ; i++ {
 	resultrune := append(resultrune, i)
 }

 return string(resultrune)
*/
}

