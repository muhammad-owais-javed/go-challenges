package sprint

func GetLastRune(s string) rune {

 li := len(s)-2
 runes := []rune(s)
 return runes[li]

}
