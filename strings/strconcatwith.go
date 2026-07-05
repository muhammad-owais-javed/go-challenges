package sprint

func StrConcatWith(strs []string, sep string) string {
	result := ""

	for i, s := range strs {
		if i > 0 {
			result += sep
		}
		result += s
	}

	return result
}