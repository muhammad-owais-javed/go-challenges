package sprint

func Countdown(n int) string {

	var result string

	for i:=n ; i > 0; i-=2 {

	numberAsChar := '0' + i
	result += string(numberAsChar)
	result += ", "

	}

	result = result[:len(result)-2]
	return result + ", 0!"

}


"6
