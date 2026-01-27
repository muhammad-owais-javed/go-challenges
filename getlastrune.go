package sprint

func GetLastRune(s string) rune {

 li := len(s)-1
 runes := []rune(s)
 return runes[li]

}
