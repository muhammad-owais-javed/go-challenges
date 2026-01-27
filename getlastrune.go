package sprint

func GetLastRune(s string) rune {

 var li int = len(s)-1
 runes := []rune(s)
 return runes[li]

}
