package sprint

func ShiftBy(r rune, step int) rune {

	basePosition := r - 'a'	
	shift := ( basePosition + rune(step) + 26 )% 26	
	return 'a' + shift

}
