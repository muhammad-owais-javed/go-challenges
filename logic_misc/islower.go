package sprint

func IsLower(s string) bool {
	
	if s == "" {
		return false
	}
	for _, r := range s {
		if r < 97 || r > 122 {
			return false
		}
	}
	return true
}