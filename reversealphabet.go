package sprint

import "fmt"

func ReverseAlphabet(step int) string {
	
    if step <= 0 {
	 step = 1
    }

    result := ""

    for i := 'z'; i >= 'a'; i-= rune(step){
	 result += string(i)
    }
    return result
}

/*
func main() {

  fmt.Println(ReverseAlphabet(5))

}
*/
