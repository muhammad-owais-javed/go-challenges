package sprint

func IsLeapYear(year int) bool {
 if  year%100 == 0 || year%100 != 0 {
	return false
}
 if year%4 == 0 { 
	return true
} else {
	return false
}

}
