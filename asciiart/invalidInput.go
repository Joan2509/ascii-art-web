package asciiart

func Unprintable(input string) bool {
	for _, char := range input {
		if !(char >= 32 && char <= 126) {
			return true
		}
	}
	return false
}
