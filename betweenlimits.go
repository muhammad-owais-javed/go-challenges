package sprint

func BetweenLimits(from, to rune) string {
 for i := from+1; i < to ; i++ {
	resultrune := append(resultrune, i)
 }
 return string(resultrune)
}

