package sprint

func IsNumeric(s string) bool {

  var check bool
  for _, r := range s {

	if  r < "0" || r > "9" {
	return false 
	}
  }

  return true
}
