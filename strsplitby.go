package sprint

func StrSplitBy(s, sep string) []string {
	var result []string
	word := ""

	for i := 0; i < len(s); {
		
		if s[i:i+len(sep)] == sep {
			result = append(result, word)
			word = ""
			i += len(sep)
		} else {
			word += string(s[i])
			i++
		}
	}

	result = append(result, word)
	return result
}
