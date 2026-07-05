package sprint
//import "fmt"


func Countdown(n int) string {	
	result := ""
	
	for i := n; i > 0;  i -= 2 {
		result += string('0' + rune(i)) + ", "
	}
//	fmt.Println(string(result))
	result = result+"0!"
	return result
}


