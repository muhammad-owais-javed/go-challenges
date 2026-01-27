package sprint

func GetLastRune(s string) rune {

 var last rune
 for _, r := range s {

	last = r
 }

   return last

}
