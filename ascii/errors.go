package ascii

func Errors(s string) (str string) {
	for _, ch := range s {
		if !(ch >= 32 && ch <= 126) {
			str = "An unprinttable character was detected in your string"
		} else {
			str += string(ch)
		}
	}
	return
}
