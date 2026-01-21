package sprint

func BetweenLimits(from, to rune) string {

 if from > to {
	from, to = to, from
 }

 var resultrune []rune
 for i:=from+1 ; i<to ; i++ {
 	resultrune := append(resultrune, i)
 }

 return string(resultrune)

}

