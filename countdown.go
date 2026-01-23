package sprint

func Countdown(n int) string {

	var result []string

	for i:=n ; i > 0; i-=2 {

	numberAsChar := '0' + i
	result = append(result, str.conv.Itoa(i))
	

	}

	mainpart := strings.Join(result, ",")
	return mainpart + ", 0!"

}

