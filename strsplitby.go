package sprint
//import "fmt"
func StrSplitBy(s, sep string) []string {
	
	var result []string //NIL slice
	part := ""

		//For empty
	if s == "" && sep == "" {
		//result = []string{}  //Empty Slice
		//fmt.Println("If Condition")
		return nil
	}

	for i := 0; i < len(s); {
	
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			result = append(result, part)
			part = ""
			i += len(sep)
		} else {
			part += string(s[i])
			i++
		}
	
	}

	result = append(result, part)
	return result

}
